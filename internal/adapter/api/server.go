package api

import (
	"context"

	connect "connectrpc.com/connect"
	mainv1 "github.com/halcyon-org/kizuna/gen/belifeline/v1"
	"github.com/halcyon-org/kizuna/gen/belifeline/v1/mainv1connect"
)

type ServerServiceHandlerImpl struct {
}

func NewServerServiceHandler() mainv1connect.ServerServiceHandler {
	return &ServerServiceHandlerImpl{}
}

const StatusResOk = "OK"

func (s *ServerServiceHandlerImpl) Health(ctx context.Context, req *connect.Request[mainv1.HealthRequest]) (*connect.Response[mainv1.HealthResponse], error) {
	res := connect.NewResponse(&mainv1.HealthResponse{Status: StatusResOk})
	return res, nil
}
