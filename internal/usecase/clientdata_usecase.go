package usecase

import (
	"context"

	entrepo "github.com/halcyon-org/kizuna/internal/adapter/repository/ent"
	"github.com/halcyon-org/kizuna/internal/domain/domain"
	"github.com/halcyon-org/kizuna/internal/util"
)

type ClientDataUsecase interface {
	CreateClientData(ctx context.Context, username string) (*domain.ClientData, error)
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
