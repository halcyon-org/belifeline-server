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

type AdminServiceHandlerImpl struct {
	clientInformationUsecase usecase.ClientInformationUsecase
}

func NewAdminServiceHandler(clientInformationUsecase usecase.ClientInformationUsecase) mainv1connect.AdminServiceHandler {
	return &AdminServiceHandlerImpl{
		clientInformationUsecase: clientInformationUsecase,
	}
}

func (s *AdminServiceHandlerImpl) ClientSet(ctx context.Context, req *connect.Request[mainv1.ClientSetRequest]) (*connect.Response[mainv1.ClientSetResponse], error) {
	// FIXME: Create and Set
	user, err := s.clientInformationUsecase.CreateClientInformation(ctx, req.Msg.Username)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	apiData := domain.ToApiClientInformation(*user)
	res := connect.NewResponse(&mainv1.ClientSetResponse{ClientInformation: &apiData})

	return res, nil
}

const limit = 100

func (s *AdminServiceHandlerImpl) ClientList(ctx context.Context, req *connect.Request[mainv1.ClientListRequest]) (*connect.Response[mainv1.ClientListResponse], error) {
	dataList, err := s.clientInformationUsecase.ListClientInformation(ctx, limit)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if len(dataList) >= limit {
		return nil, status.Error(codes.Internal, fmt.Sprintf("might have more data than %d\n", limit))
	}

	var apiDataList = make([]*v1.ClientInformation, len(dataList))
	for i, data := range dataList {
		apiData := domain.ToApiClientInformation(*data)
		apiDataList[i] = &apiData
	}
	res := connect.NewResponse(&mainv1.ClientListResponse{ClientInformationList: apiDataList})

	return res, nil
}

func (s *AdminServiceHandlerImpl) ClientDelete(ctx context.Context, req *connect.Request[mainv1.ClientDeleteRequest]) (*connect.Response[mainv1.ClientDeleteResponse], error) {
	id, err := s.clientInformationUsecase.DeleteClientInformation(ctx, req.Msg.ClientId.Value)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	res := connect.NewResponse(&mainv1.ClientDeleteResponse{ClientId: &v1.ULID{Value: id}})

	return res, nil
}

func (s *AdminServiceHandlerImpl) ClientRevoke(ctx context.Context, req *connect.Request[mainv1.ClientRevokeRequest]) (*connect.Response[mainv1.ClientRevokeResponse], error) {
	id, apiKey, err := s.clientInformationUsecase.RevokeApiKey(ctx, req.Msg.ClientId.Value)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	res := connect.NewResponse(&mainv1.ClientRevokeResponse{ClientId: &v1.ULID{Value: id}, ApiKey: &v1.ApiKey{Key: apiKey}})

	return res, nil
}

func (s *AdminServiceHandlerImpl) ExternalInformationSet(ctx context.Context, req *connect.Request[mainv1.ExternalInformationSetRequest]) (*connect.Response[mainv1.ExternalInformationSetResponse], error) {
	return nil, status.Error(codes.Unimplemented, "method ExternalInformationSet not implemented")
}

func (s *AdminServiceHandlerImpl) ExternalInformationDelete(ctx context.Context, req *connect.Request[mainv1.ExternalInformationDeleteRequest]) (*connect.Response[mainv1.ExternalInformationDeleteResponse], error) {
	return nil, status.Error(codes.Unimplemented, "method ExternalInformationDelete not implemented")
}

func (s *AdminServiceHandlerImpl) KoyoSet(ctx context.Context, req *connect.Request[mainv1.KoyoSetRequest]) (*connect.Response[mainv1.KoyoSetResponse], error) {
	return nil, status.Error(codes.Unimplemented, "method KoyoSet not implemented")
}

func (s *AdminServiceHandlerImpl) KoyoDelete(ctx context.Context, req *connect.Request[mainv1.KoyoDeleteRequest]) (*connect.Response[mainv1.KoyoDeleteResponse], error) {
	return nil, status.Error(codes.Unimplemented, "method KoyoDelete not implemented")
}

func (s *AdminServiceHandlerImpl) KoyoApiRevoke(ctx context.Context, req *connect.Request[mainv1.KoyoApiRevokeRequest]) (*connect.Response[mainv1.KoyoApiRevokeResponse], error) {
	return nil, status.Error(codes.Unimplemented, "method KoyoApiRevoke not implemented")
}
