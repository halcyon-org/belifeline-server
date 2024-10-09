package api

import (
	"context"

	connect "connectrpc.com/connect"
	mainv1 "github.com/halcyon-org/kizuna/gen/belifeline/v1"
	"github.com/halcyon-org/kizuna/gen/belifeline/v1/mainv1connect"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type KoyoServiceHandlerImpl struct {
}

func NewKoyoServiceHandler() mainv1connect.KoyoServiceHandler {
	return &KoyoServiceHandlerImpl{}
}

func (s *KoyoServiceHandlerImpl) KoyoCreate(ctx context.Context, req *connect.Request[mainv1.KoyoCreateRequest]) (*connect.Response[mainv1.KoyoCreateResponse], error) {
	return nil, status.Error(codes.Unimplemented, "method KoyoCreate not implemented")
}

func (s *KoyoServiceHandlerImpl) KoyoList(ctx context.Context, req *connect.Request[mainv1.KoyoListRequest], stream *connect.ServerStream[mainv1.KoyoListResponse]) error {
	return status.Error(codes.Unimplemented, "method KoyoList not implemented")
}

func (s *KoyoServiceHandlerImpl) KoyoDelete(ctx context.Context, req *connect.Request[mainv1.KoyoDeleteRequest]) (*connect.Response[mainv1.KoyoDeleteResponse], error) {
	return nil, status.Error(codes.Unimplemented, "method KoyoDelete not implemented")
}

func (s *KoyoServiceHandlerImpl) KoyoApiRevoke(ctx context.Context, req *connect.Request[mainv1.KoyoApiRevokeRequest]) (*connect.Response[mainv1.KoyoApiRevokeResponse], error) {
	return nil, status.Error(codes.Unimplemented, "method KoyoApiRevoke not implemented")
}

func (s *KoyoServiceHandlerImpl) KoyoDataAdd(context.Context, *connect.Request[mainv1.KoyoDataAddRequest]) (*connect.Response[mainv1.KoyoDataAddResponse], error) {
	return nil, status.Error(codes.Unimplemented, "method KoyoDataAdd not implemented")
}
