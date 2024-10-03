package api

import (
	"context"

	connect "connectrpc.com/connect"
	mainv1 "github.com/halcyon-org/kizuna/gen/belifeline/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *BeLifelineServerImpl) ClientCreate(ctx context.Context, req *connect.Request[mainv1.ClientCreateRequest]) (*connect.Response[mainv1.ClientCreateResponse], error) {
	return nil, status.Error(codes.Unimplemented, "method ClientCreate not implemented")
}

func (s *BeLifelineServerImpl) ClientList(ctx context.Context, req *connect.Request[mainv1.ClientListRequest], stream *connect.ServerStream[mainv1.ClientListResponse]) error {
	return status.Error(codes.Unimplemented, "method ClientList not implemented")
}

func (s *BeLifelineServerImpl) ClientDelete(ctx context.Context, req *connect.Request[mainv1.ClientDeleteRequest]) (*connect.Response[mainv1.ClientDeleteResponse], error) {
	return nil, status.Error(codes.Unimplemented, "method ClientDelete not implemented")
}

func (s *BeLifelineServerImpl) ClientRevoke(ctx context.Context, req *connect.Request[mainv1.ClientRevokeRequest]) (*connect.Response[mainv1.ClientRevokeResponse], error) {
	return nil, status.Error(codes.Unimplemented, "method ClientRevoke not implemented")
}
