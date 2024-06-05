package rpchandlers

import (
	"github.com/rustweave-network/rustweaved/app/appmessage"
	"github.com/rustweave-network/rustweaved/app/rpc/rpccontext"
	"github.com/rustweave-network/rustweaved/infrastructure/network/netadapter/router"
)

// HandleGetHeaders handles the respectively named RPC command
func HandleGetHeaders(context *rpccontext.Context, _ *router.Router, request appmessage.Message) (appmessage.Message, error) {
	response := &appmessage.GetHeadersResponseMessage{}
	response.Error = appmessage.RPCErrorf("not implemented")
	return response, nil
}
