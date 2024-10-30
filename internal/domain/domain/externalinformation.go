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

func ToDomainExternalInformation(e ent.ExternalInformation) ExternalInformation {
	return ExternalInformation{
		ID:                 string(e.ID),
		Name:               e.Name,
		Description:        e.Description,
		License:            e.License,
		LicenseDescription: e.LicenseDescription,
		APIKey:             e.APIKey,
		FirstEntryAt:       e.FirstEntryAt,
		LastUpdatedAt:      e.LastUpdatedAt,
	}
}

func ToAPIExternalInformation(d ExternalInformation) v1.ExternalInformation {
	return v1.ExternalInformation{
		ExternalId:          &v1.ULID{Value: d.ID},
		ExternalName:        &d.Name,
		ExternalDescription: &d.Description,
		License:             &d.License,
		LicenseDescription:  &d.LicenseDescription,
		ApiKey:              &v1.APIKey{Key: d.APIKey},
		FirstEntryAt:        timestamppb.New(d.FirstEntryAt),
		LastUpdatedAt:       timestamppb.New(d.LastUpdatedAt),
	}
}
