package domain

import (
	"time"

	"github.com/halcyon-org/kizuna/gen/ent"
)

type KoyoInformation struct {
	ID           string
	Name         string
	Description  string
	Params       map[string]string
	Scales       []float64
	Version      string
	License      string
	DataType     string
	APIKey       string
	FirstEntryAt time.Time
	LastEntryAt  time.Time
	UpdatedAt    time.Time
}

func ToDomainKoyoInformation(e ent.KoyoInformation) KoyoInformation {
	return KoyoInformation{
		ID:           string(e.ID),
		Name:         e.Name,
		Description:  e.Description,
		Params:       e.Params,
		Scales:       e.Scales,
		Version:      e.Version,
		License:      e.License,
		DataType:     string(e.DataType),
		APIKey:       e.APIKey,
		FirstEntryAt: e.FirstEntryAt,
		LastEntryAt:  e.LastEntryAt,
		UpdatedAt:    e.UpdatedAt,
	}
}
