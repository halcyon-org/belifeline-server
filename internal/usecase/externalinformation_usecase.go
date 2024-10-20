package usecase

import (
	"context"
	"errors"

	"github.com/halcyon-org/kizuna/ent/schema/pulid"
	v1 "github.com/halcyon-org/kizuna/gen/belifeline/models/v1"
	"github.com/halcyon-org/kizuna/gen/ent"
	entrepo "github.com/halcyon-org/kizuna/internal/adapter/repository/ent"
	"github.com/halcyon-org/kizuna/internal/domain/domain"

	"github.com/halcyon-org/kizuna/internal/util"
)

type ExternalInformationUsecase interface {
	SetExternalInformation(ctx context.Context, externalInformation *v1.ExternalInformation) (*domain.ExternalInformation, error)
}

type externalInformationUsecaseImpl struct {
	externalInformationRepository entrepo.ExternalInformationRepository
}

func NewExternalInformationUsecase(externalInformationRepository entrepo.ExternalInformationRepository) ExternalInformationUsecase {
	return &externalInformationUsecaseImpl{
		externalInformationRepository: externalInformationRepository,
	}
}

var ErrorPropertyNotSet = errors.New("property not set")

func (u *externalInformationUsecaseImpl) SetExternalInformation(ctx context.Context, externalInformation *v1.ExternalInformation) (*domain.ExternalInformation, error) {
	var (
		data *ent.ExternalInformation
		err  error
	)

	if externalInformation.ExternalId == nil {
		if externalInformation.ExternalName == nil || externalInformation.ExternalDescription == nil || externalInformation.License == nil || externalInformation.LicenseDescription == nil {
			return nil, ErrorPropertyNotSet
		}
		data, err = u.externalInformationRepository.CreateExternalInformation(ctx, *externalInformation.ExternalName, *externalInformation.ExternalDescription, *externalInformation.License, *externalInformation.LicenseDescription, util.GenApiKey())
	} else {
		data, err = u.externalInformationRepository.UpdateExternalInformation(ctx, pulid.ID(externalInformation.ExternalId.Value), externalInformation.ExternalName, externalInformation.ExternalDescription, externalInformation.License, externalInformation.LicenseDescription)
	}
	if err != nil {
		return nil, err
	}

	domainData := domain.ToDomainExternalInformation(*data)
	return &domainData, nil
}
