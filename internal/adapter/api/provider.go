package api

import (
	"context"
	"fmt"

	connect "connectrpc.com/connect"
	v1 "github.com/halcyon-org/kizuna/gen/belifeline/models/v1"
	mainv1 "github.com/halcyon-org/kizuna/gen/belifeline/v1"
	"github.com/halcyon-org/kizuna/gen/belifeline/v1/mainv1connect"
	"github.com/halcyon-org/kizuna/internal/domain/domain"
	"github.com/halcyon-org/kizuna/internal/usecase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProviderServiceHandlerImpl struct {
	koyoInformationUsecase usecase.KoyoInformationUsecase
}

func NewProviderServiceHandler(koyoInformationUsecase usecase.KoyoInformationUsecase) mainv1connect.ProviderServiceHandler {
	return &ProviderServiceHandlerImpl{
		koyoInformationUsecase: koyoInformationUsecase,
	}
}

func (s *ProviderServiceHandlerImpl) ExternalInformationList(ctx context.Context, req *connect.Request[mainv1.ExternalInformationListRequest]) (*connect.Response[mainv1.ExternalInformationListResponse], error) {
	return nil, status.Error(codes.Unimplemented, "method ExternalInformationList not implemented")
}

func (s *ProviderServiceHandlerImpl) KoyoList(ctx context.Context, req *connect.Request[mainv1.KoyoListRequest]) (*connect.Response[mainv1.KoyoListResponse], error) {
	dataList, err := s.koyoInformationUsecase.ListKoyoInformation(ctx, limit)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if len(dataList) >= limit {
		return nil, status.Error(codes.Internal, fmt.Sprintf("might have more data than %d\n", limit))
	}

	apiDataList := make([]*v1.KoyoInformation, len(dataList))
	for i, data := range dataList {
		apiData := domain.ToAPIKoyoInformation(*data)
		apiDataList[i] = &apiData
	}
	res := connect.NewResponse(&mainv1.KoyoListResponse{KoyoInformationList: apiDataList})

	return res, nil
}

func (s *ProviderServiceHandlerImpl) ExternalInformationGet(context.Context, *connect.Request[mainv1.ExternalInformationGetRequest]) (*connect.Response[mainv1.ExternalInformationGetResponse], error) {
	return nil, status.Error(codes.Unimplemented, "method ExternalInformationGet not implemented")
}

func (s *ProviderServiceHandlerImpl) KoyoDataGet(context.Context, *connect.Request[mainv1.KoyoDataGetRequest]) (*connect.Response[mainv1.KoyoDataGetResponse], error) {
	return nil, status.Error(codes.Unimplemented, "method KoyoDataGet not implemented")
}
