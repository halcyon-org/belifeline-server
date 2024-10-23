package api

import (
	"context"

	connect "connectrpc.com/connect"
	mainv1 "github.com/halcyon-org/kizuna/gen/belifeline/v1"
	"github.com/halcyon-org/kizuna/gen/belifeline/v1/mainv1connect"
	"github.com/halcyon-org/kizuna/internal/usecase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type KoyoServiceHandlerImpl struct {
	KoyoInformationUsecase usecase.KoyoInformationUsecase
}

func NewKoyoServiceHandler(KoyoInformationUsecase usecase.KoyoInformationUsecase) mainv1connect.KoyoServiceHandler {
	return &KoyoServiceHandlerImpl{
		KoyoInformationUsecase: KoyoInformationUsecase,
	}
}

func (s *KoyoServiceHandlerImpl) KoyoUpdate(context.Context, *connect.Request[mainv1.KoyoUpdateRequest]) (*connect.Response[mainv1.KoyoUpdateResponse], error) {
	return nil, status.Error(codes.Unimplemented, "method KoyoUpdate not implemented")
}

func (s *KoyoServiceHandlerImpl) KoyoDataAdd(context.Context, *connect.Request[mainv1.KoyoDataAddRequest]) (*connect.Response[mainv1.KoyoDataAddResponse], error) {
	return nil, status.Error(codes.Unimplemented, "method KoyoDataAdd not implemented")
}
