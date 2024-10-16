package usecase

import (
	entrepo "github.com/halcyon-org/kizuna/internal/adapter/repository/ent"
)

type ExternalInformationUsecase interface {
}

type externalInformationUsecaseImpl struct {
	externalInformationRepository entrepo.ExternalInformationRepository
}

func NewExternalInformationUsecase(externalInformationRepository entrepo.ExternalInformationRepository) ExternalInformationUsecase {
	return &externalInformationUsecaseImpl{
		externalInformationRepository: externalInformationRepository,
	}
}
