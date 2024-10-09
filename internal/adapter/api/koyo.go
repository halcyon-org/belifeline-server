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

func (s *KoyoServiceHandlerImpl) KoyoDataAdd(context.Context, *connect.Request[mainv1.KoyoDataAddRequest]) (*connect.Response[mainv1.KoyoDataAddResponse], error) {
	return nil, status.Error(codes.Unimplemented, "method KoyoDataAdd not implemented")
}
