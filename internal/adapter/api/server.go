package api

import (
	"context"

	connect "connectrpc.com/connect"
	mainv1 "github.com/halcyon-org/kizuna/gen/belifeline/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *BeLifelineServer) Status(ctx context.Context, req *connect.Request[mainv1.StatusRequest]) (*connect.Response[mainv1.StatusResponse], error) {
	return nil, status.Error(codes.Unimplemented, "method Status not implemented")
}
