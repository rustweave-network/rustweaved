package client

import (
	"context"
	"github.com/rustweave-network/rustweaved/cmd/rustweavewallet/daemon/server"
	"time"

	"github.com/pkg/errors"

	"github.com/rustweave-network/rustweaved/cmd/rustweavewallet/daemon/pb"
	"google.golang.org/grpc"
)

// Connect connects to the rustweavewalletd server, and returns the client instance
func Connect(address string) (pb.RustweavewalletdClient, func(), error) {
	// Connection is local, so 1 second timeout is sufficient
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(server.MaxDaemonSendMsgSize)))
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, nil, errors.New("rustweavewallet daemon is not running, start it with `rustweavewallet start-daemon`")
		}
		return nil, nil, err
	}

	return pb.NewRustweavewalletdClient(conn), func() {
		conn.Close()
	}, nil
}
