package api

import (
	"context"

	connect "connectrpc.com/connect"
	v1 "github.com/halcyon-org/kizuna/gen/belifeline/models/v1"
	mainv1 "github.com/halcyon-org/kizuna/gen/belifeline/v1"
	"github.com/halcyon-org/kizuna/gen/belifeline/v1/mainv1connect"
	"github.com/halcyon-org/kizuna/internal/usecase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type KoyoServiceHandlerImpl struct {
	koyoDataUsecase usecase.KoyoDataUsecase
}

func NewKoyoServiceHandler(koyoDataUsecase usecase.KoyoDataUsecase) mainv1connect.KoyoServiceHandler {
	return &KoyoServiceHandlerImpl{
		koyoDataUsecase: koyoDataUsecase,
	}
}

func (s *KoyoServiceHandlerImpl) KoyoUpdate(context.Context, *connect.Request[mainv1.KoyoUpdateRequest]) (*connect.Response[mainv1.KoyoUpdateResponse], error) {
	return nil, status.Error(codes.Unimplemented, "method KoyoUpdate not implemented")
}

func (s *KoyoServiceHandlerImpl) KoyoDataAdd(ctx context.Context, req *connect.Request[mainv1.KoyoDataAddRequest]) (*connect.Response[mainv1.KoyoDataAddResponse], error) {
	id, err := s.koyoDataUsecase.CreateKoyoData(ctx, req.Msg.KoyoData)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	res := connect.NewResponse(&mainv1.KoyoDataAddResponse{KoyoData: &v1.ULID{Value: id}})

	return res, nil
}
