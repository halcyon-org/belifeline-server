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
	clientInformationUsecase   usecase.ClientInformationUsecase
	koyoInformationUsecase     usecase.KoyoInformationUsecase
	externalInformationUsecase usecase.ExternalInformationUsecase
}

func NewAdminServiceHandler(
	clientInformationUsecase usecase.ClientInformationUsecase,
	koyoInformationUsecase usecase.KoyoInformationUsecase,
	externalInformationUsecase usecase.ExternalInformationUsecase,
) mainv1connect.AdminServiceHandler {
	return &AdminServiceHandlerImpl{
		clientInformationUsecase:   clientInformationUsecase,
		koyoInformationUsecase:     koyoInformationUsecase,
		externalInformationUsecase: externalInformationUsecase,
	}
}

func (s *AdminServiceHandlerImpl) ClientSet(
	ctx context.Context,
	req *connect.Request[mainv1.ClientSetRequest],
) (*connect.Response[mainv1.ClientSetResponse], error) {
	// FIXME: Create and Set
	user, err := s.clientInformationUsecase.CreateClientInformation(ctx, req.Msg.GetUsername())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	apiData := domain.ToAPIClientInformation(*user)
	res := connect.NewResponse(&mainv1.ClientSetResponse{ClientInformation: &apiData})

	return res, nil
}

const limit = 100

func (s *AdminServiceHandlerImpl) ClientList(
	ctx context.Context,
	_ *connect.Request[mainv1.ClientListRequest],
) (*connect.Response[mainv1.ClientListResponse], error) {
	dataList, err := s.clientInformationUsecase.ListClientInformation(ctx, limit)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if len(dataList) >= limit {
		return nil, status.Error(codes.Internal, fmt.Sprintf("might have more data than %d\n", limit))
	}

	apiDataList := make([]*v1.ClientInformation, len(dataList))

	for i, data := range dataList {
		apiData := domain.ToAPIClientInformation(*data)
		apiDataList[i] = &apiData
	}

	res := connect.NewResponse(&mainv1.ClientListResponse{ClientInformationList: apiDataList})

	return res, nil
}

func (s *AdminServiceHandlerImpl) ClientDelete(
	ctx context.Context,
	req *connect.Request[mainv1.ClientDeleteRequest],
) (*connect.Response[mainv1.ClientDeleteResponse], error) {
	clientID, err := s.clientInformationUsecase.DeleteClientInformation(ctx, req.Msg.GetClientId().GetValue())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	res := connect.NewResponse(&mainv1.ClientDeleteResponse{ClientId: &v1.ULID{Value: clientID}})

	return res, nil
}

func (s *AdminServiceHandlerImpl) ClientRevoke(
	ctx context.Context,
	req *connect.Request[mainv1.ClientRevokeRequest],
) (*connect.Response[mainv1.ClientRevokeResponse], error) {
	clientID, apiKey, err := s.clientInformationUsecase.RevokeAPIKey(ctx, req.Msg.GetClientId().GetValue())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	res := connect.NewResponse(&mainv1.ClientRevokeResponse{
		ClientId: &v1.ULID{Value: clientID},
		ApiKey:   &v1.APIKey{Key: apiKey},
	})

	return res, nil
}

func (s *AdminServiceHandlerImpl) ExternalInformationSet(
	ctx context.Context,
	req *connect.Request[mainv1.ExternalInformationSetRequest],
) (*connect.Response[mainv1.ExternalInformationSetResponse], error) {
	externalInformation := req.Msg.GetExternalInformation()
	if externalInformation.GetFirstEntryAt() != nil ||
		externalInformation.GetLastUpdatedAt() != nil ||
		externalInformation.UpdatedHistory != nil {
		return nil, status.Error(codes.InvalidArgument, NewValidationError("time should not be set").Error())
	}

	data, err := s.externalInformationUsecase.SetExternalInformation(ctx, externalInformation)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	apiData := domain.ToAPIExternalInformation(*data)
	res := connect.NewResponse(&mainv1.ExternalInformationSetResponse{ExternalInformation: &apiData})

	return res, nil
}

func (s *AdminServiceHandlerImpl) ExternalInformationDelete(
	context.Context,
	*connect.Request[mainv1.ExternalInformationDeleteRequest],
) (*connect.Response[mainv1.ExternalInformationDeleteResponse], error) {
	return nil, status.Error(codes.Unimplemented, "method ExternalInformationDelete not implemented")
}

func (s *AdminServiceHandlerImpl) KoyoCreate(
	context.Context,
	*connect.Request[mainv1.KoyoCreateRequest],
) (*connect.Response[mainv1.KoyoCreateResponse], error) {
	return nil, status.Error(codes.Unimplemented, "method KoyoSet not implemented")
}

func (s *AdminServiceHandlerImpl) KoyoDelete(
	context.Context,
	*connect.Request[mainv1.KoyoDeleteRequest],
) (*connect.Response[mainv1.KoyoDeleteResponse], error) {
	return nil, status.Error(codes.Unimplemented, "method KoyoDelete not implemented")
}

func (s *AdminServiceHandlerImpl) KoyoAPIRevoke(
	context.Context,
	*connect.Request[mainv1.KoyoAPIRevokeRequest],
) (*connect.Response[mainv1.KoyoAPIRevokeResponse], error) {
	return nil, status.Error(codes.Unimplemented, "method KoyoAPIRevoke not implemented")
}
