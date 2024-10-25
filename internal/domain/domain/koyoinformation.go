package domain

import (
	"time"

	v1 "github.com/halcyon-org/kizuna/gen/belifeline/models/v1"
	"github.com/halcyon-org/kizuna/gen/ent"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type DataType string

const (
	DataTypeUnspecified DataType = "unspecified"
	DataTypeImage       DataType = "image"
	DataTypeCsv         DataType = "csv"
	DataTypeJSON        DataType = "json"
)

type KoyoInformation struct {
	ID           string
	Name         string
	Description  string
	Params       map[string]string
	Scales       []float64
	Version      string
	License      string
	DataType     DataType
	APIKey       string
	FirstEntryAt time.Time
	LastEntryAt  time.Time
	UpdatedAt    time.Time
}

func ToDomainKoyoInformation(ent ent.KoyoInformation) KoyoInformation {
	return KoyoInformation{
		ID:           string(ent.ID),
		Name:         ent.Name,
		Description:  ent.Description,
		Params:       ent.Params,
		Scales:       ent.Scales,
		Version:      ent.Version,
		License:      ent.License,
		DataType:     DataType(ent.DataType),
		APIKey:       ent.APIKey,
		FirstEntryAt: ent.FirstEntryAt,
		LastEntryAt:  ent.LastEntryAt,
		UpdatedAt:    ent.UpdatedAt,
	}
}

func ToAPIKoyoInformation(d KoyoInformation) v1.KoyoInformation {
	scales := make([]float32, len(d.Scales))
	for i, scale := range d.Scales {
		scales[i] = float32(scale)
	}

	var dataType v1.DataType

	switch d.DataType {
	case DataTypeUnspecified:
		dataType = v1.DataType_DATA_TYPE_UNSPECIFIED
	case DataTypeImage:
		dataType = v1.DataType_DATA_TYPE_IMAGE
	case DataTypeCsv:
		dataType = v1.DataType_DATA_TYPE_CSV
	case DataTypeJSON:
		dataType = v1.DataType_DATA_TYPE_JSON
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
		ApiKey:          &v1.APIKey{Key: d.APIKey},
		FirstEntryAt:    timestamppb.New(d.FirstEntryAt),
		LastEntryAt:     timestamppb.New(d.LastEntryAt),
		LastUpdatedAt:   timestamppb.New(d.UpdatedAt),
	}
}
