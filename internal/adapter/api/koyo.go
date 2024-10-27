package api

import (
	"context"

	connect "connectrpc.com/connect"
	v1 "github.com/halcyon-org/kizuna/gen/belifeline/models/v1"
	mainv1 "github.com/halcyon-org/kizuna/gen/belifeline/v1"
	"github.com/halcyon-org/kizuna/gen/belifeline/v1/mainv1connect"
	"github.com/halcyon-org/kizuna/internal/domain/domain"
	"github.com/halcyon-org/kizuna/internal/usecase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type KoyoServiceHandlerImpl struct {
	koyoInformationUsecase usecase.KoyoInformationUsecase
	koyoDataUsecase        usecase.KoyoDataUsecase
}

func NewKoyoServiceHandler(koyoInformationUsecase usecase.KoyoInformationUsecase, koyoDataUsecase usecase.KoyoDataUsecase) mainv1connect.KoyoServiceHandler {
	return &KoyoServiceHandlerImpl{
		koyoInformationUsecase: koyoInformationUsecase,
		koyoDataUsecase:        koyoDataUsecase,
	}
}

func (s *KoyoServiceHandlerImpl) KoyoUpdate(ctx context.Context, req *connect.Request[mainv1.KoyoUpdateRequest]) (*connect.Response[mainv1.KoyoUpdateResponse], error) {
	koyoInformation := req.Msg.KoyoInformation
	if koyoInformation.FirstEntryAt != nil || koyoInformation.LastUpdatedAt != nil || koyoInformation.LastEntryAt != nil {
		return nil, status.Error(codes.InvalidArgument, NewValidationError("time should not be set").Error())
	}

	if koyoInformation.KoyoId != nil {
		return nil, status.Error(codes.InvalidArgument, NewValidationError("koyo id should not be set").Error())
	}

	data, err := s.koyoInformationUsecase.UpdateKoyoInformation(ctx, koyoInformation)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	apiData := domain.ToAPIKoyoInformation(*data)
	res := connect.NewResponse(&mainv1.KoyoUpdateResponse{KoyoInformation: &apiData})

	return res, nil
}

func (s *KoyoServiceHandlerImpl) KoyoDataAdd(ctx context.Context, req *connect.Request[mainv1.KoyoDataAddRequest]) (*connect.Response[mainv1.KoyoDataAddResponse], error) {
	id, err := s.koyoDataUsecase.CreateKoyoData(ctx, req.Msg.KoyoData)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	res := connect.NewResponse(&mainv1.KoyoDataAddResponse{KoyoData: &v1.ULID{Value: id}})

	return res, nil
}
