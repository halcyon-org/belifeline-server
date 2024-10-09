package api

import (
	"context"

	connect "connectrpc.com/connect"
	mainv1 "github.com/halcyon-org/kizuna/gen/belifeline/v1"
	"github.com/halcyon-org/kizuna/gen/belifeline/v1/mainv1connect"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ExtInfoServiceHandlerImpl struct {
}

func NewExtInfoServiceHandler() mainv1connect.ExtInfoServiceHandler {
	return &ExtInfoServiceHandlerImpl{}
}

func (s *ExtInfoServiceHandlerImpl) ExtInfoCreate(ctx context.Context, req *connect.Request[mainv1.ExtInfoCreateRequest]) (*connect.Response[mainv1.ExtInfoCreateResponse], error) {
	return nil, status.Error(codes.Unimplemented, "method ExtInfoCreate not implemented")
}

func (s *ExtInfoServiceHandlerImpl) ExtInfoList(ctx context.Context, req *connect.Request[mainv1.ExtInfoListRequest], stream *connect.ServerStream[mainv1.ExtInfoListResponse]) error {
	return status.Error(codes.Unimplemented, "method ExtInfoList not implemented")
}

func (s *ExtInfoServiceHandlerImpl) ExtInfoDelete(ctx context.Context, req *connect.Request[mainv1.ExtInfoDeleteRequest]) (*connect.Response[mainv1.ExtInfoDeleteResponse], error) {
	return nil, status.Error(codes.Unimplemented, "method ExtInfoDelete not implemented")
}

func (s *ExtInfoServiceHandlerImpl) ExtInfoUpdateNotification(context.Context, *connect.Request[mainv1.ExtInfoUpdateNotificationRequest]) (*connect.Response[mainv1.ExtInfoUpdateNotificationResponse], error) {
	return nil, status.Error(codes.Unimplemented, "method ExtInfoRevoke not implemented")
}
