package ent

import (
	"context"

	"github.com/halcyon-org/kizuna/ent"
)

type ClientDataRepository interface {
	CreateClientData(cxt context.Context, username string, apiKey string) (*ent.ClientData, error)
}

type clientDataRepositoryImpl struct {
	DB *ent.Client
}

func NewClientDataRepository(db *ent.Client) ClientDataRepository {
	return &clientDataRepositoryImpl{
		DB: db,
	}
}

func (r *clientDataRepositoryImpl) CreateClientData(ctx context.Context, username string, apiKey string) (*ent.ClientData, error) {
	clientData, err := r.DB.ClientData.Create().
		SetUsername(username).
		SetAPIKey(apiKey).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return clientData, nil
}
