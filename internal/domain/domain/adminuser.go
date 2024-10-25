package domain

import (
	"time"

	"github.com/halcyon-org/kizuna/gen/ent"
)

type AdminUser struct {
	ID            string
	Name          string
	APIKey        string
	CreatedAt     time.Time
	LastUsedAt    time.Time
	LastUpdatedAt time.Time
}

func ToDomainAdminUser(ent ent.AdminUser) AdminUser {
	return AdminUser{
		ID:            string(ent.ID),
		Name:          ent.Name,
		APIKey:        ent.APIKey,
		CreatedAt:     ent.CreatedAt,
		LastUsedAt:    ent.LastUsedAt,
		LastUpdatedAt: ent.LastUpdatedAt,
	}
}
