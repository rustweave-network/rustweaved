package server

import (
	"context"
	"github.com/rustweave-network/rustweaved/cmd/kaspawallet/daemon/pb"
	"github.com/rustweave-network/rustweaved/version"
)

func (s *server) GetVersion(_ context.Context, _ *pb.GetVersionRequest) (*pb.GetVersionResponse, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return &pb.GetVersionResponse{
		Version: version.Version(),
	}, nil
}
