package ent

import (
	"context"
	"time"

	"github.com/halcyon-org/kizuna/ent/schema/pulid"
	"github.com/halcyon-org/kizuna/gen/ent"
	"github.com/halcyon-org/kizuna/gen/ent/clientinformation"
)

type ClientInformationRepository interface {
	CreateClientInformation(cxt context.Context, username string, apiKey string) (*ent.ClientInformation, error)
	GetAllClientInformation(cxt context.Context, limit int32) ([]*ent.ClientInformation, error)
	DeleteClientInformation(ctx context.Context, client_id pulid.ID) (*pulid.ID, error)
	UpdateAPIKey(ctx context.Context, client_id pulid.ID, apiKey string) (*pulid.ID, string, error)
	GetClientInformationByAPIKey(ctx context.Context, apiKey string) (*ent.ClientInformation, error)
}

type clientInformationRepositoryImpl struct {
	DB *ent.Client
}

func NewClientInformationRepository(db *ent.Client) ClientInformationRepository {
	return &clientInformationRepositoryImpl{
		DB: db,
	}
}

func (r *clientInformationRepositoryImpl) CreateClientInformation(ctx context.Context, username string, apiKey string) (*ent.ClientInformation, error) {
	clientInformation, err := r.DB.ClientInformation.Create().
		SetUsername(username).
		SetAPIKey(apiKey).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return clientInformation, nil
}

func (r *clientInformationRepositoryImpl) GetAllClientInformation(ctx context.Context, limit int32) ([]*ent.ClientInformation, error) {
	clientInformationList, err := r.DB.ClientInformation.Query().
		Limit(int(limit)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return clientInformationList, nil
}

func (r *clientInformationRepositoryImpl) DeleteClientInformation(ctx context.Context, client_id pulid.ID) (*pulid.ID, error) {
	err := r.DB.ClientInformation.DeleteOneID(client_id).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return &client_id, nil
}

func (r *clientInformationRepositoryImpl) UpdateAPIKey(ctx context.Context, client_id pulid.ID, apiKey string) (*pulid.ID, string, error) {
	clientInformation, err := r.DB.ClientInformation.UpdateOneID(client_id).
		SetAPIKey(apiKey).
		SetLastUpdatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, "", err
	}

	return &clientInformation.ID, clientInformation.APIKey, nil
}

func (r *clientInformationRepositoryImpl) GetClientInformationByAPIKey(ctx context.Context, apiKey string) (*ent.ClientInformation, error) {
	return r.DB.ClientInformation.Query().Where(clientinformation.APIKey(apiKey)).Only(ctx)
}
