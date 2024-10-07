package api

import (
	"github.com/halcyon-org/kizuna/gen/belifeline/v1/mainv1connect"
	"github.com/halcyon-org/kizuna/internal/usecase"
)

type BeLifelineServerImpl struct {
	koyoInformationUsecase usecase.KoyoInformationUsecase
	clientDataUsecase      usecase.ClientDataUsecase
	authUsecase            usecase.AuthUsecase
}

func NewBeLifelineServiceHandler(koyoInformationUsecase usecase.KoyoInformationUsecase, clientDataUsecase usecase.ClientDataUsecase, authUsecase usecase.AuthUsecase) mainv1connect.BeLifelineServiceHandler {
	return &BeLifelineServerImpl{
		koyoInformationUsecase: koyoInformationUsecase,
		clientDataUsecase:      clientDataUsecase,
		authUsecase:            authUsecase,
	}
}
