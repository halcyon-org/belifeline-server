package api

import (
	"context"

	connect "connectrpc.com/connect"
	mainv1 "github.com/halcyon-org/kizuna/gen/belifeline/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *BeLifelineServer) ExtInfoCreate(ctx context.Context, req *connect.Request[mainv1.ExtInfoCreateRequest]) (*connect.Response[mainv1.ExtInfoCreateResponse], error) {
	return nil, status.Error(codes.Unimplemented, "method ExtInfoCreate not implemented")
}

func (s *BeLifelineServer) ExtInfoList(ctx context.Context, req *connect.Request[mainv1.ExtInfoListRequest], stream *connect.ServerStream[mainv1.ExtInfoListResponse]) error {
	return status.Error(codes.Unimplemented, "method ExtInfoList not implemented")
}

func (s *BeLifelineServer) ExtInfoDelete(ctx context.Context, req *connect.Request[mainv1.ExtInfoDeleteRequest]) (*connect.Response[mainv1.ExtInfoDeleteResponse], error) {
	return nil, status.Error(codes.Unimplemented, "method ExtInfoDelete not implemented")
}
