package server

import (
	"context"
	"time"

	"github.com/rustweave-network/rustweaved/app/appmessage"
	"github.com/rustweave-network/rustweaved/cmd/rustweavewallet/daemon/pb"
	"github.com/rustweave-network/rustweaved/cmd/rustweavewallet/librustweavewallet"
	"github.com/rustweave-network/rustweaved/cmd/rustweavewallet/librustweavewallet/serialization"
	"github.com/rustweave-network/rustweaved/domain/consensus/model/externalapi"
	"github.com/rustweave-network/rustweaved/domain/consensus/utils/consensushashing"
	"github.com/rustweave-network/rustweaved/infrastructure/network/rpcclient"
	"github.com/pkg/errors"
)

func (s *server) Broadcast(_ context.Context, request *pb.BroadcastRequest) (*pb.BroadcastResponse, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	txIDs, err := s.broadcast(request.Transactions, request.IsDomain)
	if err != nil {
		return nil, err
	}

	return &pb.BroadcastResponse{TxIDs: txIDs}, nil
}

func (s *server) broadcast(transactions [][]byte, isDomain bool) ([]string, error) {

	txIDs := make([]string, len(transactions))
	var tx *externalapi.DomainTransaction
	var err error

	for i, transaction := range transactions {

		if isDomain {
			tx, err = serialization.DeserializeDomainTransaction(transaction)
			if err != nil {
				return nil, err
			}
		} else if !isDomain { //default in proto3 is false
			tx, err = librustweavewallet.ExtractTransaction(transaction, s.keysFile.ECDSA)
			if err != nil {
				return nil, err
			}
		}

		txIDs[i], err = sendTransaction(s.rpcClient, tx)
		if err != nil {
			return nil, err
		}

		for _, input := range tx.Inputs {
			s.usedOutpoints[input.PreviousOutpoint] = time.Now()
		}
	}

	s.forceSync()
	return txIDs, nil
}

func sendTransaction(client *rpcclient.RPCClient, tx *externalapi.DomainTransaction) (string, error) {
	submitTransactionResponse, err := client.SubmitTransaction(appmessage.DomainTransactionToRPCTransaction(tx), consensushashing.TransactionID(tx).String(), false)
	if err != nil {
		return "", errors.Wrapf(err, "error submitting transaction")
	}
	return submitTransactionResponse.TransactionID, nil
}
