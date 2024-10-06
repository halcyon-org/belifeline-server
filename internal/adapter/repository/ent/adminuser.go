package ent

import (
	"context"

	"github.com/halcyon-org/kizuna/ent"
	"github.com/halcyon-org/kizuna/ent/adminuser"
)

type AdminUserRepository interface {
	GetAdminUserByAPIKey(ctx context.Context, apiKey string) (*ent.AdminUser, error)
}

type adminUserRepositoryImpl struct {
	DB *ent.Client
}

func NewAdminUserRepository(db *ent.Client) AdminUserRepository {
	return &adminUserRepositoryImpl{
		DB: db,
	}
}

func (r *adminUserRepositoryImpl) GetAdminUserByAPIKey(ctx context.Context, apiKey string) (*ent.AdminUser, error) {
	return r.DB.AdminUser.Query().Where(adminuser.APIKey(apiKey)).Only(ctx)
}
