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

type ProviderServiceHandlerImpl struct {
	clientDataUsecase usecase.ClientDataUsecase
}

func NewProviderServiceHandler(clientDataUsecase usecase.ClientDataUsecase) mainv1connect.ProviderServiceHandler {
	return &ProviderServiceHandlerImpl{
		clientDataUsecase: clientDataUsecase,
	}
}

func (s *ProviderServiceHandlerImpl) ClientCreate(ctx context.Context, req *connect.Request[mainv1.ClientCreateRequest]) (*connect.Response[mainv1.ClientCreateResponse], error) {
	user, err := s.clientDataUsecase.CreateClientData(ctx, req.Msg.Username)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	apiData := domain.ToApiClientData(*user)
	res := connect.NewResponse(&mainv1.ClientCreateResponse{ClientData: &apiData})

	return res, nil
}

func (s *ProviderServiceHandlerImpl) ClientList(ctx context.Context, req *connect.Request[mainv1.ClientListRequest], stream *connect.ServerStream[mainv1.ClientListResponse]) error {
	return status.Error(codes.Unimplemented, "method ClientList not implemented")
}

func (s *ProviderServiceHandlerImpl) ClientDelete(ctx context.Context, req *connect.Request[mainv1.ClientDeleteRequest]) (*connect.Response[mainv1.ClientDeleteResponse], error) {
	return nil, status.Error(codes.Unimplemented, "method ClientDelete not implemented")
}

func (s *ProviderServiceHandlerImpl) ClientRevoke(ctx context.Context, req *connect.Request[mainv1.ClientRevokeRequest]) (*connect.Response[mainv1.ClientRevokeResponse], error) {
	return nil, status.Error(codes.Unimplemented, "method ClientRevoke not implemented")
}
