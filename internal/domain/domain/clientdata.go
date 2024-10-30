package domain

import (
	"time"

	v1 "github.com/halcyon-org/kizuna/gen/belifeline/models/v1"
	"github.com/halcyon-org/kizuna/gen/ent"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ClientInformation struct {
	ID            string
	Username      string
	APIKey        string
	CreatedAt     time.Time
	LastUsedAt    time.Time
	LastUpdatedAt time.Time
}

func ToDomainClientInformation(e ent.ClientInformation) ClientInformation {
	return ClientInformation{
		ID:            string(e.ID),
		Username:      e.Username,
		APIKey:        e.APIKey,
		CreatedAt:     e.CreatedAt,
		LastUsedAt:    e.LastUsedAt,
		LastUpdatedAt: e.LastUpdatedAt,
	}
}

func ToAPIClientInformation(d ClientInformation) v1.ClientInformation {
	return v1.ClientInformation{
		ClientId:      &v1.ULID{Value: d.ID},
		Username:      &d.Username,
		ApiKey:        &v1.APIKey{Key: d.APIKey},
		CreatedAt:     timestamppb.New(d.CreatedAt),
		LastUsedAt:    timestamppb.New(d.LastUsedAt),
		LastUpdatedAt: timestamppb.New(d.LastUpdatedAt),
	}
}
