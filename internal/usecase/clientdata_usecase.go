package usecase

import (
	"context"

	"github.com/halcyon-org/kizuna/ent/schema/pulid"
	entrepo "github.com/halcyon-org/kizuna/internal/adapter/repository/ent"
	"github.com/halcyon-org/kizuna/internal/domain/domain"
	"github.com/halcyon-org/kizuna/internal/util"
)

type ClientInformationUsecase interface {
	CreateClientInformation(ctx context.Context, username string) (*domain.ClientInformation, error)
	ListClientInformation(ctx context.Context, limit int32) ([]*domain.ClientInformation, error)
	DeleteClientInformation(ctx context.Context, clientID string) (string, error)
	RevokeAPIKey(ctx context.Context, clientID string) (string, string, error)
}

type clientInformationUsecaseImpl struct {
	clientInformationRepository entrepo.ClientInformationRepository
}

func NewClientInformationUsecase(
	clientInformationRepository entrepo.ClientInformationRepository,
) ClientInformationUsecase {
	return &clientInformationUsecaseImpl{
		clientInformationRepository: clientInformationRepository,
	}
}

func (u *clientInformationUsecaseImpl) CreateClientInformation(
	ctx context.Context,
	username string,
) (*domain.ClientInformation, error) {
	data, err := u.clientInformationRepository.CreateClientInformation(ctx, username, util.GenAPIKey())
	if err != nil {
		return nil, err
	}

	domainData := domain.ToDomainClientInformation(*data)

	return &domainData, nil
}

func (u *clientInformationUsecaseImpl) ListClientInformation(
	ctx context.Context,
	limit int32,
) ([]*domain.ClientInformation, error) {
	dataList, err := u.clientInformationRepository.GetAllClientInformation(ctx, limit)
	if err != nil {
		return nil, err
	}

	domainDataList := make([]*domain.ClientInformation, len(dataList))

	for i, data := range dataList {
		domainData := domain.ToDomainClientInformation(*data)
		domainDataList[i] = &domainData
	}

	return domainDataList, nil
}

func (u *clientInformationUsecaseImpl) DeleteClientInformation(ctx context.Context, clientID string) (string, error) {
	id, err := u.clientInformationRepository.DeleteClientInformation(ctx, pulid.ID(clientID))
	if err != nil {
		return "", err
	}

	return string(*id), nil
}

func (u *clientInformationUsecaseImpl) RevokeAPIKey(ctx context.Context, clientID string) (string, string, error) {
	id, apiKey, err := u.clientInformationRepository.UpdateAPIKey(ctx, pulid.ID(clientID), util.GenAPIKey())
	if err != nil {
		return "", "", err
	}

	return string(*id), apiKey, nil
}
