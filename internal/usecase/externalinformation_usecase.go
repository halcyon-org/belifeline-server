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
	SetExternalInformation(
		ctx context.Context,
		externalInformation *v1.ExternalInformation,
	) (*domain.ExternalInformation, error)
}

type externalInformationUsecaseImpl struct {
	externalInformationRepository entrepo.ExternalInformationRepository
}

func NewExternalInformationUsecase(
	externalInformationRepository entrepo.ExternalInformationRepository,
) ExternalInformationUsecase {
	return &externalInformationUsecaseImpl{
		externalInformationRepository: externalInformationRepository,
	}
}

var ErrPropertyNotSet = errors.New("property not set")

func (u *externalInformationUsecaseImpl) SetExternalInformation(
	ctx context.Context,
	externalInformation *v1.ExternalInformation,
) (*domain.ExternalInformation, error) {
	var (
		data *ent.ExternalInformation
		err  error
	)

	if externalInformation.GetExternalId() == nil {
		if externalInformation.ExternalName == nil ||
			externalInformation.ExternalDescription == nil ||
			externalInformation.License == nil ||
			externalInformation.LicenseDescription == nil {
			return nil, ErrPropertyNotSet
		}

		data, err = u.externalInformationRepository.CreateExternalInformation(
			ctx,
			externalInformation.GetExternalName(),
			externalInformation.GetExternalDescription(),
			externalInformation.GetLicense(),
			externalInformation.GetLicenseDescription(),
			util.GenAPIKey(),
		)
	} else {
		name := externalInformation.GetExternalName()
		description := externalInformation.GetExternalDescription()
		license := externalInformation.GetLicense()
		licenseDescription := externalInformation.GetLicenseDescription()
		data, err = u.externalInformationRepository.UpdateExternalInformation(
			ctx,
			pulid.ID(externalInformation.GetExternalId().GetValue()),
			&name,
			&description,
			&license,
			&licenseDescription,
		)
	}

	if err != nil {
		return nil, err
	}

	domainData := domain.ToDomainExternalInformation(*data)

	return &domainData, nil
}
