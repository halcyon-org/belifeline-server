package usecase

import (
	"context"
	"errors"

	"github.com/halcyon-org/kizuna/ent/schema/pulid"
	v1 "github.com/halcyon-org/kizuna/gen/belifeline/models/v1"
	entrepo "github.com/halcyon-org/kizuna/internal/adapter/repository/ent"
	"github.com/halcyon-org/kizuna/internal/domain/domain"

	"github.com/halcyon-org/kizuna/internal/util"
)

type ExternalInformationUsecase interface {
	SetExternalInformation(ctx context.Context, externalInformation *v1.ExternalInformation) (*domain.ExternalInformation, error)
	GetExternalInformation(ctx context.Context, id string) (*domain.ExternalInformation, error)
}

type externalInformationUsecaseImpl struct {
	externalInformationRepository entrepo.ExternalInformationRepository
}

func NewExternalInformationUsecase(externalInformationRepository entrepo.ExternalInformationRepository) ExternalInformationUsecase {
	return &externalInformationUsecaseImpl{
		externalInformationRepository: externalInformationRepository,
	}
}

func (u *externalInformationUsecaseImpl) SetExternalInformation(ctx context.Context, externalInformation *v1.ExternalInformation) (*domain.ExternalInformation, error) {
	if externalInformation.ExternalId == nil {
		if externalInformation.ExternalName == nil || externalInformation.ExternalDescription == nil || externalInformation.License == nil || externalInformation.LicenseDescription == nil {
			return nil, errors.New("property not set")
		}
		data, err := u.externalInformationRepository.CreateExternalInformation(ctx, *externalInformation.ExternalName, *externalInformation.ExternalDescription, *externalInformation.License, *externalInformation.LicenseDescription, util.GenApiKey())
		if err != nil {
			return nil, err
		}

		domainData := domain.ToDomainExternalInformation(*data)
		return &domainData, nil
	}

	information, err := u.externalInformationRepository.GetExternalInformationByID(ctx, pulid.ID(externalInformation.ExternalId.Value))
	if err != nil {
		return nil, err
	}

	name := information.Name
	description := information.Description
	license := information.License
	licenseDescription := information.LicenseDescription

	if externalInformation.ExternalName != nil {
		name = *externalInformation.ExternalName
	}
	if externalInformation.ExternalDescription != nil {
		description = *externalInformation.ExternalDescription
	}
	if externalInformation.License != nil {
		license = *externalInformation.License
	}
	if externalInformation.LicenseDescription != nil {
		licenseDescription = *externalInformation.LicenseDescription
	}

	data, err := u.externalInformationRepository.UpdateExternalInformation(ctx, pulid.ID(information.ID), name, description, license, licenseDescription)
	if err != nil {
		return nil, err
	}

	domainData := domain.ToDomainExternalInformation(*data)
	return &domainData, nil
}

func (u *externalInformationUsecaseImpl) GetExternalInformation(ctx context.Context, id string) (*domain.ExternalInformation, error) {
	data, err := u.externalInformationRepository.GetExternalInformationByID(ctx, pulid.ID(id))
	if err != nil {
		return nil, err
	}

	domainData := domain.ToDomainExternalInformation(*data)
	return &domainData, nil
}
