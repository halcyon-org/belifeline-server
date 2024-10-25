package ent

import (
	"context"
	"time"

	"github.com/halcyon-org/kizuna/ent/schema/pulid"
	"github.com/halcyon-org/kizuna/gen/ent"
	"github.com/halcyon-org/kizuna/gen/ent/externalinformation"
)

type ExternalInformationRepository interface {
	CreateExternalInformation(
		ctx context.Context,
		name string,
		description string,
		license string,
		licenseDescription string,
		apiKey string,
	) (*ent.ExternalInformation, error)
	UpdateExternalInformation(
		ctx context.Context,
		externalID pulid.ID,
		name *string,
		description *string,
		license *string,
		licenseDescription *string,
	) (*ent.ExternalInformation, error)
	GetExternalInformationByAPIKey(ctx context.Context, apiKey string) (*ent.ExternalInformation, error)
}

type externalInformationRepositoryImpl struct {
	DB *ent.Client
}

func NewExternalInformationRepository(db *ent.Client) ExternalInformationRepository {
	return &externalInformationRepositoryImpl{
		DB: db,
	}
}

func (r *externalInformationRepositoryImpl) CreateExternalInformation(
	ctx context.Context,
	name string,
	description string,
	license string,
	licenseDescription string,
	apiKey string,
) (*ent.ExternalInformation, error) {
	externalInformation, err := r.DB.ExternalInformation.Create().
		SetName(name).
		SetDescription(description).
		SetLicense(license).
		SetLicenseDescription(licenseDescription).
		SetAPIKey(apiKey).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return externalInformation, nil
}

func (r *externalInformationRepositoryImpl) UpdateExternalInformation(
	ctx context.Context,
	externalID pulid.ID,
	name *string,
	description *string,
	license *string,
	licenseDescription *string,
) (*ent.ExternalInformation, error) {
	update := r.DB.ExternalInformation.UpdateOneID(externalID)
	if name != nil {
		update.SetName(*name)
	}

	if description != nil {
		update.SetDescription(*description)
	}

	if license != nil {
		update.SetLicense(*license)
	}

	if licenseDescription != nil {
		update.SetLicenseDescription(*licenseDescription)
	}

	externalInformation, err := update.SetLastUpdatedAt(time.Now()).Save(ctx)
	if err != nil {
		return nil, err
	}

	return externalInformation, nil
}

func (r *externalInformationRepositoryImpl) GetExternalInformationByAPIKey(
	ctx context.Context,
	apiKey string,
) (*ent.ExternalInformation, error) {
	return r.DB.ExternalInformation.Query().Where(externalinformation.APIKey(apiKey)).Only(ctx)
}
