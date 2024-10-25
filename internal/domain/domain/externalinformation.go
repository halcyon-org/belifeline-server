package domain

import (
	"time"

	v1 "github.com/halcyon-org/kizuna/gen/belifeline/models/v1"
	"github.com/halcyon-org/kizuna/gen/ent"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ExternalInformation struct {
	ID                 string
	Name               string
	Description        string
	License            string
	LicenseDescription string
	APIKey             string
	FirstEntryAt       time.Time
	LastUpdatedAt      time.Time
}

func ToDomainExternalInformation(ent ent.ExternalInformation) ExternalInformation {
	return ExternalInformation{
		ID:                 string(ent.ID),
		Name:               ent.Name,
		Description:        ent.Description,
		License:            ent.License,
		LicenseDescription: ent.LicenseDescription,
		APIKey:             ent.APIKey,
		FirstEntryAt:       ent.FirstEntryAt,
		LastUpdatedAt:      ent.LastUpdatedAt,
	}
}

func ToAPIExternalInformation(domain ExternalInformation) v1.ExternalInformation {
	return v1.ExternalInformation{
		ExternalId:          &v1.ULID{Value: domain.ID},
		ExternalName:        &domain.Name,
		ExternalDescription: &domain.Description,
		License:             &domain.License,
		LicenseDescription:  &domain.LicenseDescription,
		ApiKey:              &v1.APIKey{Key: domain.APIKey},
		FirstEntryAt:        timestamppb.New(domain.FirstEntryAt),
		LastUpdatedAt:       timestamppb.New(domain.LastUpdatedAt),
	}
}
