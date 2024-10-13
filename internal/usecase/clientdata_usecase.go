package usecase

import (
	"context"

	"github.com/halcyon-org/kizuna/ent/schema/pulid"
	entrepo "github.com/halcyon-org/kizuna/internal/adapter/repository/ent"
	"github.com/halcyon-org/kizuna/internal/domain/domain"
	"github.com/halcyon-org/kizuna/internal/util"
)

type ClientDataUsecase interface {
	CreateClientData(ctx context.Context, username string) (*domain.ClientData, error)
	ListClientData(ctx context.Context, limit int32) ([]*domain.ClientData, error)
	DeleteClientData(ctx context.Context, client_id string) (string, error)
	RevokeApiKey(ctx context.Context, client_id string) (string, string, error)
}

type clientDataUsecaseImpl struct {
	clientDataRepository entrepo.ClientDataRepository
}

func NewClientDataUsecase(clientDataRepository entrepo.ClientDataRepository) ClientDataUsecase {
	return &clientDataUsecaseImpl{
		clientDataRepository: clientDataRepository,
	}
}

func (u *clientDataUsecaseImpl) CreateClientData(ctx context.Context, username string) (*domain.ClientData, error) {
	data, err := u.clientDataRepository.CreateClientData(ctx, username, util.GenApiKey())
	if err != nil {
		return nil, err
	}

	domainData := domain.ToDomainClientData(*data)
	return &domainData, nil
}

func (u *clientDataUsecaseImpl) ListClientData(ctx context.Context, limit int32) ([]*domain.ClientData, error) {
	dataList, err := u.clientDataRepository.GetAllClientData(ctx, limit)
	if err != nil {
		return nil, err
	}

	var domainDataList = make([]*domain.ClientData, len(dataList))
	for i, data := range dataList {
		domainData := domain.ToDomainClientData(*data)
		domainDataList[i] = &domainData
	}
	return domainDataList, nil
}

func (u *clientDataUsecaseImpl) DeleteClientData(ctx context.Context, client_id string) (string, error) {
	id, err := u.clientDataRepository.DeleteClientData(ctx, pulid.ID(client_id))
	if err != nil {
		return "", err
	}

	return string(*id), nil
}

func (u *clientDataUsecaseImpl) RevokeApiKey(ctx context.Context, client_id string) (string, string, error) {
	id, apiKey, err := u.clientDataRepository.RevokeAPIKey(ctx, pulid.ID(client_id))
	if err != nil {
		return "", "", err
	}

	return string(*id), apiKey, nil
}
