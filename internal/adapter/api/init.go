package api

import (
	"github.com/halcyon-org/kizuna/gen/belifeline/v1/mainv1connect"
	"github.com/halcyon-org/kizuna/internal/usecase"
)

type BeLifelineServerImpl struct {
	koyoInfomationUsecase usecase.KoyoInfomationUsecase
}

func NewBeLifelineServiceHandler(koyoInfomationUsecase usecase.KoyoInfomationUsecase) mainv1connect.BeLifelineServiceHandler {
	return &BeLifelineServerImpl{
		koyoInfomationUsecase: koyoInfomationUsecase,
	}
}
