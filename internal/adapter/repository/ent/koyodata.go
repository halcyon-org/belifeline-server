package ent

import (
	"context"
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/halcyon-org/kizuna/ent/schema/pulid"
	"github.com/halcyon-org/kizuna/gen/ent"
	"github.com/halcyon-org/kizuna/internal/domain/domain"
	"github.com/halcyon-org/kizuna/internal/infrastructure/infra"
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
	DB   *ent.Client
	File infra.FilesInterface
}

func NewKoyoDataRepository(db *ent.Client, file infra.FilesInterface) KoyoDataRepository {
	return &koyoDataRepositoryImpl{
		DB:   db,
		File: file,
	}
}

func (r *koyoDataRepositoryImpl) CreateKoyoData(
	ctx context.Context,
	koyoId pulid.ID,
	scale float64,
	params map[string]string,
	values []domain.Value,
) (*pulid.ID, error) {
	data, err := r.DB.KoyoData.Create().
		SetKoyoID(koyoId).
		SetScale(scale).
		SetParams(params).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	wd, err := r.File.GetWD()
	if err != nil {
		return nil, err
	}
	path := filepath.Join(wd, string(koyoId), fmt.Sprintf("%s.json", string(data.ID)))

	content, err := json.Marshal(domain.Content{ID: data.ID, Values: values})
	if err != nil {
		return nil, err
	}

	if err := r.File.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return nil, err
	}

	if err := r.File.WriteFile(path, content, 0644); err != nil {
		return nil, err
	}

	return &data.ID, nil
}
