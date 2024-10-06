package ent

import "github.com/halcyon-org/kizuna/ent"

type AdminUserRepository interface {
}

type adminUserRepositoryImpl struct {
	DB *ent.Client
}

func NewAdminUserRepository(db *ent.Client) AdminUserRepository {
	return &adminUserRepositoryImpl{
		DB: db,
	}
}
