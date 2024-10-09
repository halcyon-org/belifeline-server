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

func (s *ExtInfoServiceHandlerImpl) ExtInfoUpdateNotification(context.Context, *connect.Request[mainv1.ExtInfoUpdateNotificationRequest]) (*connect.Response[mainv1.ExtInfoUpdateNotificationResponse], error) {
	return nil, status.Error(codes.Unimplemented, "method ExtInfoRevoke not implemented")
}
