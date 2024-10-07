package usecase

import "github.com/halcyon-org/kizuna/internal/adapter/repository/ent"

type KoyoInformationUsecase interface {
}

type koyoInformationUsecaseImpl struct {
	koyoInformationRepository ent.KoyoInformationRepository
}

func NewKoyoInformationUsecase(koyoInformationRepository ent.KoyoInformationRepository) KoyoInformationUsecase {
	return &koyoInformationUsecaseImpl{
		koyoInformationRepository: koyoInformationRepository,
	}
}
