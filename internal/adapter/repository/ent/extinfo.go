package ent

import (
	"context"

	"github.com/halcyon-org/kizuna/ent"
	"github.com/halcyon-org/kizuna/ent/externalinformation"
)

type ExternalInformationRepository interface {
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

func (r *externalInformationRepositoryImpl) GetExternalInformationByAPIKey(ctx context.Context, apiKey string) (*ent.ExternalInformation, error) {
	return r.DB.ExternalInformation.Query().Where(externalinformation.APIKey(apiKey)).Only(ctx)
}
