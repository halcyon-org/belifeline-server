package domain

import (
	"time"

	"github.com/halcyon-org/kizuna/ent"
)

type AdminUser struct {
	ID            string
	Name          string
	APIKey        string
	CreatedAt     time.Time
	LastUsedAt    time.Time
	LastUpdatedAt time.Time
}

func ToDomainAdminUser(e ent.AdminUser) AdminUser {
	return AdminUser{
		ID:            string(e.ID),
		Name:          e.Name,
		APIKey:        e.APIKey,
		CreatedAt:     e.CreatedAt,
		LastUsedAt:    e.LastUsedAt,
		LastUpdatedAt: e.LastUpdatedAt,
	}
}
