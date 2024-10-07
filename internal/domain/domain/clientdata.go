package domain

import (
	"time"

	"github.com/halcyon-org/kizuna/ent"
	v1 "github.com/halcyon-org/kizuna/gen/belifeline/models/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ClientData struct {
	ID            string
	Username      string
	APIKey        string
	CreatedAt     time.Time
	LastUsedAt    time.Time
	LastUpdatedAt time.Time
}

func ToDomainClientData(e ent.ClientData) ClientData {
	return ClientData{
		ID:            string(e.ID),
		Username:      e.Username,
		APIKey:        e.APIKey,
		CreatedAt:     e.CreatedAt,
		LastUsedAt:    e.LastUsedAt,
		LastUpdatedAt: e.LastUpdatedAt,
	}
}

func ToApiClientData(d ClientData) v1.ClientData {
	return v1.ClientData{
		ClientId:      &v1.ULID{Value: d.ID},
		Username:      d.Username,
		ApiKey:        &v1.ApiKey{Key: d.APIKey},
		CreatedAt:     timestamppb.New(d.CreatedAt),
		LastUsedAt:    timestamppb.New(d.LastUsedAt),
		LastUpdatedAt: timestamppb.New(d.LastUpdatedAt),
	}
}
