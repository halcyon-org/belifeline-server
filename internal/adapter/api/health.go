package api

import (
	"context"

	connect "connectrpc.com/connect"
	mainv1 "github.com/halcyon-org/kizuna/gen/belifeline/v1"
	"github.com/halcyon-org/kizuna/gen/belifeline/v1/mainv1connect"
)

type HealthServiceHandlerImpl struct {
}

func NewHealthServiceHandler() mainv1connect.HealthServiceHandler {
	return &HealthServiceHandlerImpl{}
}

func (s *HealthServiceHandlerImpl) Check(ctx context.Context, req *connect.Request[mainv1.CheckRequest]) (*connect.Response[mainv1.CheckResponse], error) {
	res := connect.NewResponse(&mainv1.CheckResponse{Status: mainv1.ServingStatus_SERVING_STATUS_SERVING})
	return res, nil
}

func (s *HealthServiceHandlerImpl) Watch(ctx context.Context, req *connect.Request[mainv1.WatchRequest], stream *connect.ServerStream[mainv1.WatchResponse]) error {
	return nil
}
