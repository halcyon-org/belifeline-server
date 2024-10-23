package ent

import (
	"context"

	"github.com/halcyon-org/kizuna/ent/schema/pulid"
	"github.com/halcyon-org/kizuna/gen/ent"
	"github.com/halcyon-org/kizuna/gen/ent/koyoinformation"
)

type KoyoInformationRepository interface {
	CreateKoyoInformation(ctx context.Context, name string, description string, external []pulid.ID, params map[string]string, scales []float64, dataIds []pulid.ID, version string, license string, dataType koyoinformation.DataType, apiKey string) (*ent.KoyoInformation, error)
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

func (r *koyoInformationRepositoryImpl) GetKoyoInformationByAPIKey(ctx context.Context, apiKey string) (*ent.KoyoInformation, error) {
	return r.DB.KoyoInformation.Query().Where(koyoinformation.APIKey(apiKey)).Only(ctx)
}
