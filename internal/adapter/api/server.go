package api

import (
	"context"

	connect "connectrpc.com/connect"
	mainv1 "github.com/halcyon-org/kizuna/gen/belifeline/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *BeLifelineServer) Health(ctx context.Context, req *connect.Request[mainv1.HealthRequest]) (*connect.Response[mainv1.HealthResponse], error) {
	return nil, status.Error(codes.Unimplemented, "method Health not implemented")
}
