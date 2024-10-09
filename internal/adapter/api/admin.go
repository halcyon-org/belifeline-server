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

type AdminServiceHandlerImpl struct {
	clientDataUsecase usecase.ClientDataUsecase
}

func NewAdminServiceHandler(clientDataUsecase usecase.ClientDataUsecase) mainv1connect.AdminServiceHandler {
	return &AdminServiceHandlerImpl{
		clientDataUsecase: clientDataUsecase,
	}
}

func (s *AdminServiceHandlerImpl) ClientSet(ctx context.Context, req *connect.Request[mainv1.ClientSetRequest]) (*connect.Response[mainv1.ClientSetResponse], error) {
	// FIXME: Create and Set
	user, err := s.clientDataUsecase.CreateClientData(ctx, req.Msg.Username)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	apiData := domain.ToApiClientData(*user)
	res := connect.NewResponse(&mainv1.ClientSetResponse{ClientData: &apiData})

	return res, nil
}

func (s *AdminServiceHandlerImpl) ClientDelete(ctx context.Context, req *connect.Request[mainv1.ClientDeleteRequest]) (*connect.Response[mainv1.ClientDeleteResponse], error) {
	return nil, status.Error(codes.Unimplemented, "method ClientDelete not implemented")
}

func (s *AdminServiceHandlerImpl) ClientRevoke(ctx context.Context, req *connect.Request[mainv1.ClientRevokeRequest]) (*connect.Response[mainv1.ClientRevokeResponse], error) {
	return nil, status.Error(codes.Unimplemented, "method ClientRevoke not implemented")
}

func (s *AdminServiceHandlerImpl) ExtInfoSet(ctx context.Context, req *connect.Request[mainv1.ExtInfoSetRequest]) (*connect.Response[mainv1.ExtInfoSetResponse], error) {
	return nil, status.Error(codes.Unimplemented, "method ExtInfoSet not implemented")
}

func (s *AdminServiceHandlerImpl) ExtInfoDelete(ctx context.Context, req *connect.Request[mainv1.ExtInfoDeleteRequest]) (*connect.Response[mainv1.ExtInfoDeleteResponse], error) {
	return nil, status.Error(codes.Unimplemented, "method ExtInfoDelete not implemented")
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
