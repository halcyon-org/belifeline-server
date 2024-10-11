package domain

import (
	"time"

	"github.com/halcyon-org/kizuna/ent"
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
