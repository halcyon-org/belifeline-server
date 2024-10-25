package usecase

import (
	"context"
	"errors"

	"github.com/halcyon-org/kizuna/internal/adapter/repository/ent"
	"github.com/halcyon-org/kizuna/internal/domain/domain"
)

type AuthUsecase interface {
	AuthAdminUser(ctx context.Context, apiKey string) (*domain.AdminUser, error)
	AuthClientInformation(ctx context.Context, apiKey string) (*domain.ClientInformation, error)
	AuthKoyoInformation(ctx context.Context, apiKey string) (*domain.KoyoInformation, error)
	AuthExternalInformation(ctx context.Context, apiKey string) (*domain.ExternalInformation, error)
}

type authUsecaseImpl struct {
	adminUserRepository           ent.AdminUserRepository
	clientInformationRepository   ent.ClientInformationRepository
	koyoInformationRepository     ent.KoyoInformationRepository
	externalInformationRepository ent.ExternalInformationRepository
}

func NewAuthUsecase(adminUserRepository ent.AdminUserRepository,
	clientInformationRepository ent.ClientInformationRepository,
	koyoInformationRepository ent.KoyoInformationRepository,
	externalInformationRepository ent.ExternalInformationRepository,
) AuthUsecase {
	return &authUsecaseImpl{
		adminUserRepository:           adminUserRepository,
		clientInformationRepository:   clientInformationRepository,
		koyoInformationRepository:     koyoInformationRepository,
		externalInformationRepository: externalInformationRepository,
	}
}

var ErrAuthenticationFailed = errors.New("authentication failed")

func (u *authUsecaseImpl) AuthAdminUser(ctx context.Context, apiKey string) (*domain.AdminUser, error) {
	user, err := u.adminUserRepository.GetAdminUserByAPIKey(ctx, apiKey)
	if err != nil {
		return nil, ErrAuthenticationFailed
	}

	adminUser := domain.ToDomainAdminUser(*user)

	return &adminUser, nil
}

func (u *authUsecaseImpl) AuthClientInformation(ctx context.Context, apiKey string) (*domain.ClientInformation, error) {
	data, err := u.clientInformationRepository.GetClientInformationByAPIKey(ctx, apiKey)
	if err != nil {
		return nil, ErrAuthenticationFailed
	}

	clientInformation := domain.ToDomainClientInformation(*data)

	return &clientInformation, nil
}

func (u *authUsecaseImpl) AuthKoyoInformation(ctx context.Context, apiKey string) (*domain.KoyoInformation, error) {
	info, err := u.koyoInformationRepository.GetKoyoInformationByAPIKey(ctx, apiKey)
	if err != nil {
		return nil, ErrAuthenticationFailed
	}

	koyoInformation := domain.ToDomainKoyoInformation(*info)

	return &koyoInformation, nil
}

func (u *authUsecaseImpl) AuthExternalInformation(
	ctx context.Context,
	apiKey string,
) (*domain.ExternalInformation, error) {
	info, err := u.externalInformationRepository.GetExternalInformationByAPIKey(ctx, apiKey)
	if err != nil {
		return nil, ErrAuthenticationFailed
	}

	externalInformation := domain.ToDomainExternalInformation(*info)

	return &externalInformation, nil
}
