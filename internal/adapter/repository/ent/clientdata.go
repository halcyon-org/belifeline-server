package ent

import (
	"context"

	"github.com/halcyon-org/kizuna/ent"
	"github.com/halcyon-org/kizuna/ent/clientdata"
	"github.com/halcyon-org/kizuna/ent/schema/pulid"
)

type ClientDataRepository interface {
	CreateClientData(cxt context.Context, username string, apiKey string) (*ent.ClientData, error)
	ListClientData(cxt context.Context, limit int32) ([]*ent.ClientData, error)
	DeleteClientData(ctx context.Context, client_id pulid.ID) (*pulid.ID, error)
	GetClientDataByAPIKey(ctx context.Context, apiKey string) (*ent.ClientData, error)
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

func (r *clientDataRepositoryImpl) ListClientData(ctx context.Context, limit int32) ([]*ent.ClientData, error) {
	clientDataList, err := r.DB.ClientData.Query().
		Limit(int(limit)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return clientDataList, nil
}

func (r *clientDataRepositoryImpl) DeleteClientData(ctx context.Context, client_id pulid.ID) (*pulid.ID, error) {
	err := r.DB.ClientData.DeleteOneID(client_id).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return &client_id, nil
}

func (r *clientDataRepositoryImpl) GetClientDataByAPIKey(ctx context.Context, apiKey string) (*ent.ClientData, error) {
	return r.DB.ClientData.Query().Where(clientdata.APIKey(apiKey)).Only(ctx)
}
