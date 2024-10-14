package api

import (
	"context"

	connect "connectrpc.com/connect"
	mainv1 "github.com/halcyon-org/kizuna/gen/belifeline/v1"
	"github.com/halcyon-org/kizuna/gen/belifeline/v1/mainv1connect"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ExternalInformationServiceHandlerImpl struct {
}

func NewExternalInformationServiceHandler() mainv1connect.ExternalInformationServiceHandler {
	return &ExternalInformationServiceHandlerImpl{}
}

func (s *ExternalInformationServiceHandlerImpl) ExternalInformationUpdateNotification(context.Context, *connect.Request[mainv1.ExternalInformationUpdateNotificationRequest]) (*connect.Response[mainv1.ExternalInformationUpdateNotificationResponse], error) {
	return nil, status.Error(codes.Unimplemented, "method ExternalInformationRevoke not implemented")
}
