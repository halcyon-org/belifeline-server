package ent

import (
	"context"

	"github.com/halcyon-org/kizuna/gen/ent"
	"github.com/halcyon-org/kizuna/gen/ent/koyoinformation"
)

type KoyoInformationRepository interface {
	GetKoyoInformationByAPIKey(ctx context.Context, apiKey string) (*ent.KoyoInformation, error)
}

type koyoInformationRepositoryImpl struct {
	DB *ent.Client
}

func NewKoyoInformationRepository(db *ent.Client) KoyoInformationRepository {
	return &koyoInformationRepositoryImpl{
		DB: db,
	}
}

func (r *koyoInformationRepositoryImpl) GetKoyoInformationByAPIKey(ctx context.Context, apiKey string) (*ent.KoyoInformation, error) {
	return r.DB.KoyoInformation.Query().Where(koyoinformation.APIKey(apiKey)).Only(ctx)
}
