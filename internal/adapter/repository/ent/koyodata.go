package ent

import (
	"context"

	"github.com/halcyon-org/kizuna/ent/schema/pulid"
	"github.com/halcyon-org/kizuna/gen/ent"
)

type KoyoDataRepository interface {
	CreateKoyoData(ctx context.Context, koyoId pulid.ID, scale float64, params map[string]string) (*pulid.ID, error)
}

type koyoDataRepositoryImpl struct {
	DB *ent.Client
}

func NewKoyoDataRepository(db *ent.Client) KoyoDataRepository {
	return &koyoDataRepositoryImpl{
		DB: db,
	}
}

func (r *koyoDataRepositoryImpl) CreateKoyoData(ctx context.Context, koyoId pulid.ID, scale float64, params map[string]string) (*pulid.ID, error) {
	data, err := r.DB.KoyoData.Create().
		SetKoyoID(koyoId).
		SetScale(scale).
		SetParams(params).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &data.ID, nil
}
