package domain

import (
	"time"

	v1 "github.com/halcyon-org/kizuna/gen/belifeline/models/v1"
	"github.com/halcyon-org/kizuna/gen/ent"
	"google.golang.org/protobuf/types/known/timestamppb"
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

func ToAPIKoyoInformation(d KoyoInformation) v1.KoyoInformation {
	scales := make([]float32, len(d.Scales))
	for i, scale := range d.Scales {
		scales[i] = float32(scale)
	}

	var dataType v1.DataType
	switch d.DataType {
	case "image":
		dataType = v1.DataType_DATA_TYPE_IMAGE
	case "csv":
		dataType = v1.DataType_DATA_TYPE_CSV
	case "json":
		dataType = v1.DataType_DATA_TYPE_JSON
	default:
		dataType = v1.DataType_DATA_TYPE_UNSPECIFIED
	}

	return v1.KoyoInformation{
		KoyoId:          &v1.ULID{Value: d.ID},
		KoyoName:        &d.Name,
		KoyoDescription: &d.Description,
		KoyoParams:      d.Params,
		KoyoScales:      scales,
		Version:         &v1.Version{Value: d.Version},
		License:         &d.License,
		DataType:        &dataType,
		ApiKey:          &v1.ApiKey{Key: d.APIKey},
		FirstEntryAt:    timestamppb.New(d.FirstEntryAt),
		LastEntryAt:     timestamppb.New(d.LastEntryAt),
		LastUpdatedAt:   timestamppb.New(d.UpdatedAt),
	}
}
