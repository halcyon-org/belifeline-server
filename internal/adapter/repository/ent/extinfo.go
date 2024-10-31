package ent

import (
	"context"
	"time"

	"github.com/halcyon-org/kizuna/ent/schema/pulid"
	"github.com/halcyon-org/kizuna/gen/ent"
	"github.com/halcyon-org/kizuna/gen/ent/externalinformation"
)

type ExternalInformationRepository interface {
	CreateExternalInformation(ctx context.Context, name string, description string, license string, licenseDescription string, apiKey string) (*ent.ExternalInformation, error)
	UpdateExternalInformation(ctx context.Context, external_id pulid.ID, name *string, description *string, license *string, licenseDescription *string) (*ent.ExternalInformation, error)
	DeleteExternalInformation(ctx context.Context, externalId pulid.ID) (*pulid.ID, error)
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

func (r *externalInformationRepositoryImpl) CreateExternalInformation(ctx context.Context, name string, description string, license string, licenseDescription string, apiKey string) (*ent.ExternalInformation, error) {
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

func (r *externalInformationRepositoryImpl) UpdateExternalInformation(ctx context.Context, external_id pulid.ID, name *string, description *string, license *string, licenseDescription *string) (*ent.ExternalInformation, error) {
	u := r.DB.ExternalInformation.UpdateOneID(external_id)
	if name != nil {
		u.SetName(*name)
	}
	if description != nil {
		u.SetDescription(*description)
	}
	if license != nil {
		u.SetLicense(*license)
	}
	if licenseDescription != nil {
		u.SetLicenseDescription(*licenseDescription)
	}
	externalInformation, err := u.SetLastUpdatedAt(time.Now()).Save(ctx)
	if err != nil {
		return nil, err
	}
	return externalInformation, nil
}

func (r *externalInformationRepositoryImpl) DeleteExternalInformation(ctx context.Context, externalId pulid.ID) (*pulid.ID, error) {
	err := r.DB.ExternalInformation.DeleteOneID(externalId).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return &externalId, nil
}

func (r *externalInformationRepositoryImpl) GetExternalInformationByAPIKey(ctx context.Context, apiKey string) (*ent.ExternalInformation, error) {
	return r.DB.ExternalInformation.Query().Where(externalinformation.APIKey(apiKey)).Only(ctx)
}
