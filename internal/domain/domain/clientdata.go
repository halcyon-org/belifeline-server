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

func ToDomainClientInformation(ent ent.ClientInformation) ClientInformation {
	return ClientInformation{
		ID:            string(ent.ID),
		Username:      ent.Username,
		APIKey:        ent.APIKey,
		CreatedAt:     ent.CreatedAt,
		LastUsedAt:    ent.LastUsedAt,
		LastUpdatedAt: ent.LastUpdatedAt,
	}
}

func ToAPIClientInformation(domain ClientInformation) v1.ClientInformation {
	return v1.ClientInformation{
		ClientId:      &v1.ULID{Value: domain.ID},
		Username:      &domain.Username,
		ApiKey:        &v1.APIKey{Key: domain.APIKey},
		CreatedAt:     timestamppb.New(domain.CreatedAt),
		LastUsedAt:    timestamppb.New(domain.LastUsedAt),
		LastUpdatedAt: timestamppb.New(domain.LastUpdatedAt),
	}
}
