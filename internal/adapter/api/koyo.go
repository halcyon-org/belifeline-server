package api

import (
	"context"

	connect "connectrpc.com/connect"
	mainv1 "github.com/halcyon-org/kizuna/gen/belifeline/v1"
	"github.com/halcyon-org/kizuna/gen/belifeline/v1/mainv1connect"
	"github.com/halcyon-org/kizuna/internal/domain/domain"
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

func (s *KoyoServiceHandlerImpl) KoyoUpdate(ctx context.Context, req *connect.Request[mainv1.KoyoUpdateRequest]) (*connect.Response[mainv1.KoyoUpdateResponse], error) {
	koyoInformation := req.Msg.KoyoInformation
	if koyoInformation.FirstEntryAt != nil || koyoInformation.LastUpdatedAt != nil || koyoInformation.LastEntryAt != nil {
		return nil, status.Error(codes.InvalidArgument, NewValidationError("time should not be set").Error())
	}

	data, err := s.KoyoInformationUsecase.UpdateKoyoInformation(ctx, koyoInformation)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	apiData := domain.ToAPIKoyoInformation(*data)
	res := connect.NewResponse(&mainv1.KoyoUpdateResponse{KoyoInformation: &apiData})

	return res, nil
}

func (s *KoyoServiceHandlerImpl) KoyoDataAdd(context.Context, *connect.Request[mainv1.KoyoDataAddRequest]) (*connect.Response[mainv1.KoyoDataAddResponse], error) {
	return nil, status.Error(codes.Unimplemented, "method KoyoDataAdd not implemented")
}
