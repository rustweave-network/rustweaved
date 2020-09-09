package blockdag

import (
	"fmt"
	"github.com/kaspanet/go-secp256k1"
	"github.com/kaspanet/kaspad/util"
	"github.com/kaspanet/kaspad/util/daghash"
	"github.com/pkg/errors"
)

// TxAcceptanceData stores a transaction together with an indication
// if it was accepted or not by some block
type TxAcceptanceData struct {
	Tx         *util.Tx
	IsAccepted bool
}

// BlockTxsAcceptanceData stores all transactions in a block with an indication
// if they were accepted or not by some other block
type BlockTxsAcceptanceData struct {
	BlockHash        daghash.Hash
	TxAcceptanceData []TxAcceptanceData
}

// MultiBlockTxsAcceptanceData stores data about which transactions were accepted by a block
// It's a slice of the block's blues block IDs and their transaction acceptance data
type MultiBlockTxsAcceptanceData []BlockTxsAcceptanceData

// FindAcceptanceData finds the BlockTxsAcceptanceData that matches blockHash
func (data MultiBlockTxsAcceptanceData) FindAcceptanceData(blockHash *daghash.Hash) (*BlockTxsAcceptanceData, bool) {
	for _, acceptanceData := range data {
		if acceptanceData.BlockHash.IsEqual(blockHash) {
			return &acceptanceData, true
		}
	}
	return nil, false
}

// TxsAcceptedByVirtual retrieves transactions accepted by the current virtual block
//
// This function MUST be called with the DAG read-lock held
func (dag *BlockDAG) TxsAcceptedByVirtual() (MultiBlockTxsAcceptanceData, error) {
	_, _, txsAcceptanceData, err := dag.pastUTXO(&dag.virtual.blockNode)
	return txsAcceptanceData, err
}

// TxsAcceptedByBlockHash retrieves transactions accepted by the given block
//
// This function MUST be called with the DAG read-lock held
func (dag *BlockDAG) TxsAcceptedByBlockHash(blockHash *daghash.Hash) (MultiBlockTxsAcceptanceData, error) {
	node, ok := dag.index.LookupNode(blockHash)
	if !ok {
		return nil, errors.Errorf("Couldn't find block %s", blockHash)
	}
	_, _, txsAcceptanceData, err := dag.pastUTXO(node)
	return txsAcceptanceData, err
}

func (dag *BlockDAG) meldVirtualUTXO(newVirtualUTXODiffSet *DiffUTXOSet) error {
	return newVirtualUTXODiffSet.meldToBase()
}

// checkDoubleSpendsWithBlockPast checks that each block transaction
// has a corresponding UTXO in the block pastUTXO.
func checkDoubleSpendsWithBlockPast(pastUTXO UTXOSet, blockTransactions []*util.Tx) error {
	for _, tx := range blockTransactions {
		if tx.IsCoinBase() {
			continue
		}

		for _, txIn := range tx.MsgTx().TxIn {
			if _, ok := pastUTXO.Get(txIn.PreviousOutpoint); !ok {
				return ruleError(ErrMissingTxOut, fmt.Sprintf("missing transaction "+
					"output %s in the utxo set", txIn.PreviousOutpoint))
			}
		}
	}

	return nil
}

// verifyAndBuildUTXO verifies all transactions in the given block and builds its UTXO
// to save extra traversals it returns the transactions acceptance data, the compactFeeData
// for the new block and its multiset.
func (node *blockNode) verifyAndBuildUTXO(dag *BlockDAG, transactions []*util.Tx, fastAdd bool) (
	newBlockUTXO UTXOSet, txsAcceptanceData MultiBlockTxsAcceptanceData, newBlockFeeData compactFeeData, multiset *secp256k1.MultiSet, err error) {

	pastUTXO, selectedParentPastUTXO, txsAcceptanceData, err := dag.pastUTXO(node)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	err = node.validateAcceptedIDMerkleRoot(dag, txsAcceptanceData)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	feeData, err := dag.checkConnectToPastUTXO(node, pastUTXO, transactions, fastAdd)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	multiset, err = node.calcMultiset(dag, txsAcceptanceData, selectedParentPastUTXO)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	err = node.validateUTXOCommitment(multiset)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return pastUTXO, txsAcceptanceData, feeData, multiset, nil
}

func genesisPastUTXO(virtual *virtualBlock) UTXOSet {
	// The genesis has no past UTXO, so we create an empty UTXO
	// set by creating a diff UTXO set with the virtual UTXO
	// set, and adding all of its entries in toRemove
	diff := NewUTXODiff()
	for outpoint, entry := range virtual.utxoSet.utxoCollection {
		diff.toRemove[outpoint] = entry
	}
	genesisPastUTXO := UTXOSet(NewDiffUTXOSet(virtual.utxoSet, diff))
	return genesisPastUTXO
}

// applyBlueBlocks adds all transactions in the blue blocks to the selectedParent's past UTXO set
// Purposefully ignoring failures - these are just unaccepted transactions
// Writing down which transactions were accepted or not in txsAcceptanceData
func (node *blockNode) applyBlueBlocks(selectedParentPastUTXO UTXOSet, blueBlocks []*util.Block) (
	pastUTXO UTXOSet, multiBlockTxsAcceptanceData MultiBlockTxsAcceptanceData, err error) {

	pastUTXO = selectedParentPastUTXO.(*DiffUTXOSet).cloneWithoutBase()
	multiBlockTxsAcceptanceData = make(MultiBlockTxsAcceptanceData, len(blueBlocks))

	// Add blueBlocks to multiBlockTxsAcceptanceData in topological order. This
	// is so that anyone who iterates over it would process blocks (and transactions)
	// in their order of appearance in the DAG.
	for i := 0; i < len(blueBlocks); i++ {
		blueBlock := blueBlocks[i]
		transactions := blueBlock.Transactions()
		blockTxsAcceptanceData := BlockTxsAcceptanceData{
			BlockHash:        *blueBlock.Hash(),
			TxAcceptanceData: make([]TxAcceptanceData, len(transactions)),
		}
		isSelectedParent := i == 0

		for j, tx := range blueBlock.Transactions() {
			var isAccepted bool

			// Coinbase transaction outputs are added to the UTXO
			// only if they are in the selected parent chain.
			if !isSelectedParent && tx.IsCoinBase() {
				isAccepted = false
			} else {
				isAccepted, err = pastUTXO.AddTx(tx.MsgTx(), node.blueScore)
				if err != nil {
					return nil, nil, err
				}
			}
			blockTxsAcceptanceData.TxAcceptanceData[j] = TxAcceptanceData{Tx: tx, IsAccepted: isAccepted}
		}
		multiBlockTxsAcceptanceData[i] = blockTxsAcceptanceData
	}

	return pastUTXO, multiBlockTxsAcceptanceData, nil
}

// pastUTXO returns the UTXO of a given block's past
// To save traversals over the blue blocks, it also returns the transaction acceptance data for
// all blue blocks
func (dag *BlockDAG) pastUTXO(node *blockNode) (
	pastUTXO, selectedParentPastUTXO UTXOSet, bluesTxsAcceptanceData MultiBlockTxsAcceptanceData, err error) {

	if node.isGenesis() {
		return genesisPastUTXO(dag.virtual), nil, MultiBlockTxsAcceptanceData{}, nil
	}

	selectedParentPastUTXO, err = dag.restorePastUTXO(node.selectedParent)
	if err != nil {
		return nil, nil, nil, err
	}

	blueBlocks, err := dag.fetchBlueBlocks(node)
	if err != nil {
		return nil, nil, nil, err
	}

	pastUTXO, bluesTxsAcceptanceData, err = node.applyBlueBlocks(selectedParentPastUTXO, blueBlocks)
	if err != nil {
		return nil, nil, nil, err
	}

	return pastUTXO, selectedParentPastUTXO, bluesTxsAcceptanceData, nil
}

// restorePastUTXO restores the UTXO of a given block from its diff
func (dag *BlockDAG) restorePastUTXO(node *blockNode) (UTXOSet, error) {
	stack := []*blockNode{}

	// Iterate over the chain of diff-childs from node till virtual and add them
	// all into a stack
	for current := node; current != nil; {
		stack = append(stack, current)
		var err error
		current, err = dag.utxoDiffStore.diffChildByNode(current)
		if err != nil {
			return nil, err
		}
	}

	// Start with the top item in the stack, going over it top-to-bottom,
	// applying the UTXO-diff one-by-one.
	topNode, stack := stack[len(stack)-1], stack[:len(stack)-1] // pop the top item in the stack
	topNodeDiff, err := dag.utxoDiffStore.diffByNode(topNode)
	if err != nil {
		return nil, err
	}
	accumulatedDiff := topNodeDiff.clone()

	for i := len(stack) - 1; i >= 0; i-- {
		diff, err := dag.utxoDiffStore.diffByNode(stack[i])
		if err != nil {
			return nil, err
		}
		// Use withDiffInPlace, otherwise copying the diffs again and again create a polynomial overhead
		err = accumulatedDiff.withDiffInPlace(diff)
		if err != nil {
			return nil, err
		}
	}

	return NewDiffUTXOSet(dag.virtual.utxoSet, accumulatedDiff), nil
}

// updateTipsUTXO builds and applies new diff UTXOs for all the DAG's tips
func updateTipsUTXO(dag *BlockDAG, virtualUTXO UTXOSet) error {
	for tip := range dag.virtual.parents {
		tipPastUTXO, err := dag.restorePastUTXO(tip)
		if err != nil {
			return err
		}
		diff, err := virtualUTXO.diffFrom(tipPastUTXO)
		if err != nil {
			return err
		}
		err = dag.utxoDiffStore.setBlockDiff(tip, diff)
		if err != nil {
			return err
		}
	}

	return nil
}

// updateParents adds this block to the children sets of its parents
// and updates the diff of any parent whose DiffChild is this block
func (node *blockNode) updateParents(dag *BlockDAG, newBlockUTXO UTXOSet) error {
	node.updateParentsChildren()
	return node.updateParentsDiffs(dag, newBlockUTXO)
}

// updateParentsDiffs updates the diff of any parent whose DiffChild is this block
func (node *blockNode) updateParentsDiffs(dag *BlockDAG, newBlockUTXO UTXOSet) error {
	virtualDiffFromNewBlock, err := dag.virtual.utxoSet.diffFrom(newBlockUTXO)
	if err != nil {
		return err
	}

	err = dag.utxoDiffStore.setBlockDiff(node, virtualDiffFromNewBlock)
	if err != nil {
		return err
	}

	for parent := range node.parents {
		diffChild, err := dag.utxoDiffStore.diffChildByNode(parent)
		if err != nil {
			return err
		}
		if diffChild == nil {
			parentPastUTXO, err := dag.restorePastUTXO(parent)
			if err != nil {
				return err
			}
			err = dag.utxoDiffStore.setBlockDiffChild(parent, node)
			if err != nil {
				return err
			}
			diff, err := newBlockUTXO.diffFrom(parentPastUTXO)
			if err != nil {
				return err
			}
			err = dag.utxoDiffStore.setBlockDiff(parent, diff)
			if err != nil {
				return err
			}
		}
	}

	return nil
}