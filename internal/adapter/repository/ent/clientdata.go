package ent

import (
	"context"
	"time"

	"github.com/halcyon-org/kizuna/ent"
	"github.com/halcyon-org/kizuna/ent/clientdata"
	"github.com/halcyon-org/kizuna/ent/schema/pulid"
)

type ClientDataRepository interface {
	CreateClientData(cxt context.Context, username string, apiKey string) (*ent.ClientData, error)
	GetAllClientData(cxt context.Context, limit int32) ([]*ent.ClientData, error)
	DeleteClientData(ctx context.Context, client_id pulid.ID) (*pulid.ID, error)
	UpdateAPIKey(ctx context.Context, client_id pulid.ID, apiKey string) (*pulid.ID, string, error)
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

func (r *clientDataRepositoryImpl) GetAllClientData(ctx context.Context, limit int32) ([]*ent.ClientData, error) {
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

func (r *clientDataRepositoryImpl) UpdateAPIKey(ctx context.Context, client_id pulid.ID, apiKey string) (*pulid.ID, string, error) {
	clientData, err := r.DB.ClientData.UpdateOneID(client_id).
		SetAPIKey(apiKey).
		SetLastUpdatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, "", err
	}

	return &clientData.ID, clientData.APIKey, nil
}

func (r *clientDataRepositoryImpl) GetClientDataByAPIKey(ctx context.Context, apiKey string) (*ent.ClientData, error) {
	return r.DB.ClientData.Query().Where(clientdata.APIKey(apiKey)).Only(ctx)
}
