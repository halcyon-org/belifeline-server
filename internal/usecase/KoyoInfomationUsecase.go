package usecase

import "github.com/halcyon-org/kizuna/internal/adapter/repository/ent"

type KoyoInfomationUsecase interface {
}

type koyoInfomationUsecaseImpl struct {
	koyoInfomationRepository ent.KoyoInfomationRepository
}

func NewKoyoInfomationUsecase(koyoInfomationRepository ent.KoyoInfomationRepository) KoyoInfomationUsecase {
	return &koyoInfomationUsecaseImpl{
		koyoInfomationRepository: koyoInfomationRepository,
	}
}
