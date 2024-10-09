package api

import (
	"context"

	connect "connectrpc.com/connect"
	mainv1 "github.com/halcyon-org/kizuna/gen/belifeline/v1"
	"github.com/halcyon-org/kizuna/gen/belifeline/v1/mainv1connect"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProviderServiceHandlerImpl struct {
}

func NewProviderServiceHandler() mainv1connect.ProviderServiceHandler {
	return &ProviderServiceHandlerImpl{}
}

func (s *ProviderServiceHandlerImpl) ClientList(ctx context.Context, req *connect.Request[mainv1.ClientListRequest], stream *connect.ServerStream[mainv1.ClientListResponse]) error {
	return status.Error(codes.Unimplemented, "method ClientList not implemented")
}

func (s *ProviderServiceHandlerImpl) ExtInfoList(ctx context.Context, req *connect.Request[mainv1.ExtInfoListRequest], stream *connect.ServerStream[mainv1.ExtInfoListResponse]) error {
	return status.Error(codes.Unimplemented, "method ExtInfoList not implemented")
}

func (s *ProviderServiceHandlerImpl) KoyoList(ctx context.Context, req *connect.Request[mainv1.KoyoListRequest], stream *connect.ServerStream[mainv1.KoyoListResponse]) error {
	return status.Error(codes.Unimplemented, "method KoyoList not implemented")
}

func (s *ProviderServiceHandlerImpl) ExtInfoGet(context.Context, *connect.Request[mainv1.ExtInfoGetRequest]) (*connect.Response[mainv1.ExtInfoGetResponse], error) {
	return nil, status.Error(codes.Unimplemented, "method ExtInfoGet not implemented")
}

func (s *ProviderServiceHandlerImpl) KoyoDataGet(context.Context, *connect.Request[mainv1.KoyoDataGetRequest]) (*connect.Response[mainv1.KoyoDataGetResponse], error) {
	return nil, status.Error(codes.Unimplemented, "method KoyoDataGet not implemented")
}
