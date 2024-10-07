package usecase

import (
	"context"
	"errors"

	"github.com/halcyon-org/kizuna/internal/adapter/repository/ent"
	"github.com/halcyon-org/kizuna/internal/domain/domain"
)

type AuthUsecase interface {
	AuthAdminUser(ctx context.Context, apiKey string) (*domain.AdminUser, error)
}

type authUsecaseImpl struct {
	adminUserRepository ent.AdminUserRepository
}

func NewAuthUsecase(adminUserRepository ent.AdminUserRepository) AuthUsecase {
	return &authUsecaseImpl{
		adminUserRepository: adminUserRepository,
	}
}

var ErrorAuthenticationFailed = errors.New("Authentication failed")

func (u *authUsecaseImpl) AuthAdminUser(ctx context.Context, apiKey string) (*domain.AdminUser, error) {
	user, err := u.adminUserRepository.GetAdminUserByAPIKey(ctx, apiKey)
	if err != nil {
		return nil, ErrorAuthenticationFailed
	}
	adminUser := domain.ToDomainAdminUser(*user)
	return &adminUser, nil
}
