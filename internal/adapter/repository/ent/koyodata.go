package ent

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/halcyon-org/kizuna/ent/schema/pulid"
	"github.com/halcyon-org/kizuna/gen/ent"
	"github.com/halcyon-org/kizuna/internal/domain/domain"
)

type KoyoDataRepository interface {
	CreateKoyoData(
		ctx context.Context,
		koyoId pulid.ID,
		scale float64,
		params map[string]string,
		value []domain.Value,
	) (*pulid.ID, error)
}

type koyoDataRepositoryImpl struct {
	DB *ent.Client
}

func NewKoyoDataRepository(db *ent.Client) KoyoDataRepository {
	return &koyoDataRepositoryImpl{
		DB: db,
	}
}

func (r *koyoDataRepositoryImpl) CreateKoyoData(
	ctx context.Context,
	koyoId pulid.ID,
	scale float64,
	params map[string]string,
	value []domain.Value,
) (*pulid.ID, error) {
	data, err := r.DB.KoyoData.Create().
		SetKoyoID(koyoId).
		SetScale(scale).
		SetParams(params).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	path, err := filepath.Abs(fmt.Sprintf("./%s/%s.json", koyoId, data.ID))
	if err != nil {
		return nil, err
	}

	content, err := json.Marshal(domain.Content{ID: data.ID, Values: value})
	if err != nil {
		return nil, err
	}

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return nil, err
	}

	if err := os.WriteFile(path, content, 0644); err != nil {
		return nil, err
	}

	return &data.ID, nil
}
