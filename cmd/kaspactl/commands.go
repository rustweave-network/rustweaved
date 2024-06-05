package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/rustweave-network/rustweaved/infrastructure/network/netadapter/server/grpcserver/protowire"
)

var commandTypes = []reflect.Type{
	reflect.TypeOf(protowire.RustweavedMessage_AddPeerRequest{}),
	reflect.TypeOf(protowire.RustweavedMessage_GetConnectedPeerInfoRequest{}),
	reflect.TypeOf(protowire.RustweavedMessage_GetPeerAddressesRequest{}),
	reflect.TypeOf(protowire.RustweavedMessage_GetCurrentNetworkRequest{}),
	reflect.TypeOf(protowire.RustweavedMessage_GetInfoRequest{}),

	reflect.TypeOf(protowire.RustweavedMessage_GetBlockRequest{}),
	reflect.TypeOf(protowire.RustweavedMessage_GetBlocksRequest{}),
	reflect.TypeOf(protowire.RustweavedMessage_GetHeadersRequest{}),
	reflect.TypeOf(protowire.RustweavedMessage_GetBlockCountRequest{}),
	reflect.TypeOf(protowire.RustweavedMessage_GetBlockDagInfoRequest{}),
	reflect.TypeOf(protowire.RustweavedMessage_GetSelectedTipHashRequest{}),
	reflect.TypeOf(protowire.RustweavedMessage_GetVirtualSelectedParentBlueScoreRequest{}),
	reflect.TypeOf(protowire.RustweavedMessage_GetVirtualSelectedParentChainFromBlockRequest{}),
	reflect.TypeOf(protowire.RustweavedMessage_ResolveFinalityConflictRequest{}),
	reflect.TypeOf(protowire.RustweavedMessage_EstimateNetworkHashesPerSecondRequest{}),

	reflect.TypeOf(protowire.RustweavedMessage_GetBlockTemplateRequest{}),
	reflect.TypeOf(protowire.RustweavedMessage_SubmitBlockRequest{}),

	reflect.TypeOf(protowire.RustweavedMessage_GetMempoolEntryRequest{}),
	reflect.TypeOf(protowire.RustweavedMessage_GetMempoolEntriesRequest{}),
	reflect.TypeOf(protowire.RustweavedMessage_GetMempoolEntriesByAddressesRequest{}),

	reflect.TypeOf(protowire.RustweavedMessage_SubmitTransactionRequest{}),

	reflect.TypeOf(protowire.RustweavedMessage_GetUtxosByAddressesRequest{}),
	reflect.TypeOf(protowire.RustweavedMessage_GetBalanceByAddressRequest{}),
	reflect.TypeOf(protowire.RustweavedMessage_GetCoinSupplyRequest{}),

	reflect.TypeOf(protowire.RustweavedMessage_BanRequest{}),
	reflect.TypeOf(protowire.RustweavedMessage_UnbanRequest{}),
}

type commandDescription struct {
	name       string
	parameters []*parameterDescription
	typeof     reflect.Type
}

type parameterDescription struct {
	name   string
	typeof reflect.Type
}

func commandDescriptions() []*commandDescription {
	commandDescriptions := make([]*commandDescription, len(commandTypes))

	for i, commandTypeWrapped := range commandTypes {
		commandType := unwrapCommandType(commandTypeWrapped)

		name := strings.TrimSuffix(commandType.Name(), "RequestMessage")
		numFields := commandType.NumField()

		var parameters []*parameterDescription
		for i := 0; i < numFields; i++ {
			field := commandType.Field(i)

			if !isFieldExported(field) {
				continue
			}

			parameters = append(parameters, &parameterDescription{
				name:   field.Name,
				typeof: field.Type,
			})
		}
		commandDescriptions[i] = &commandDescription{
			name:       name,
			parameters: parameters,
			typeof:     commandTypeWrapped,
		}
	}

	return commandDescriptions
}

func (cd *commandDescription) help() string {
	sb := &strings.Builder{}
	sb.WriteString(cd.name)
	for _, parameter := range cd.parameters {
		_, _ = fmt.Fprintf(sb, " [%s]", parameter.name)
	}
	return sb.String()
}
