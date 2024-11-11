package ent

import (
	"context"
	"time"

	"github.com/halcyon-org/kizuna/ent/schema/pulid"
	"github.com/halcyon-org/kizuna/gen/ent"
	"github.com/halcyon-org/kizuna/gen/ent/koyoinformation"
)

type KoyoInformationRepository interface {
	CreateKoyoInformation(ctx context.Context, name string, description string, external []pulid.ID, params map[string]string, scales []float64, dataIds []pulid.ID, version string, license string, dataType koyoinformation.DataType, apiKey string) (*ent.KoyoInformation, error)
	UpdateKoyoInformation(ctx context.Context, koyoId pulid.ID, name *string, description *string, external *[]pulid.ID, params *map[string]string, scales *[]float64, dataIds *[]pulid.ID, version *string, license *string, dataType *koyoinformation.DataType) (*ent.KoyoInformation, error)
	DeleteKoyoInformation(ctx context.Context, koyoId pulid.ID) (*pulid.ID, error)
	GetKoyoInformationByAPIKey(ctx context.Context, apiKey string) (*ent.KoyoInformation, error)
}

type koyoInformationRepositoryImpl struct {
	DB *ent.Client
}

func NewKoyoInformationRepository(db *ent.Client) KoyoInformationRepository {
	return &koyoInformationRepositoryImpl{
		DB: db,
	}
}

func (r *koyoInformationRepositoryImpl) CreateKoyoInformation(ctx context.Context, name string, description string, external []pulid.ID, params map[string]string, scales []float64, dataIds []pulid.ID, version string, license string, dataType koyoinformation.DataType, apiKey string) (*ent.KoyoInformation, error) {
	return r.DB.KoyoInformation.Create().
		SetName(name).
		SetDescription(description).
		AddExternalIDs(external...).
		SetParams(params).
		SetScales(scales).
		AddDatumIDs(dataIds...).
		SetVersion(version).
		SetLicense(license).
		SetDataType(dataType).
		SetAPIKey(apiKey).
		Save(ctx)
}

func (r *koyoInformationRepositoryImpl) UpdateKoyoInformation(ctx context.Context, koyoId pulid.ID, name *string, description *string, external *[]pulid.ID, params *map[string]string, scales *[]float64, dataIds *[]pulid.ID, version *string, license *string, dataType *koyoinformation.DataType) (*ent.KoyoInformation, error) {
	u := r.DB.KoyoInformation.UpdateOneID(koyoId)
	if name != nil {
		u.SetName(*name)
	}
	if description != nil {
		u.SetDescription(*description)
	}
	if external != nil {
		u.AddExternalIDs(*external...)
	}
	if params != nil {
		u.SetParams(*params)
	}
	if scales != nil {
		u.SetScales(*scales)
	}
	if dataIds != nil {
		u.AddDatumIDs(*dataIds...)
	}
	if version != nil {
		u.SetVersion(*version)
	}
	if license != nil {
		u.SetLicense(*license)
	}
	if dataType != nil {
		u.SetDataType(*dataType)
	}

	return u.SetLastEntryAt(time.Now()).Save(ctx)
}

func (r *koyoInformationRepositoryImpl) DeleteKoyoInformation(ctx context.Context, koyoId pulid.ID) (*pulid.ID, error) {
	err := r.DB.KoyoInformation.DeleteOneID(koyoId).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return &koyoId, nil
}

func (r *koyoInformationRepositoryImpl) GetKoyoInformationByAPIKey(ctx context.Context, apiKey string) (*ent.KoyoInformation, error) {
	return r.DB.KoyoInformation.Query().Where(koyoinformation.APIKey(apiKey)).Only(ctx)
}
