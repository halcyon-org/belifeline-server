package usecase

import (
	"context"

	"github.com/halcyon-org/kizuna/ent/schema/pulid"
	v1 "github.com/halcyon-org/kizuna/gen/belifeline/models/v1"
	entrepo "github.com/halcyon-org/kizuna/internal/adapter/repository/ent"
	"github.com/halcyon-org/kizuna/internal/domain/domain"
)

type KoyoDataUsecase interface {
	CreateKoyoData(ctx context.Context, koyoData *v1.KoyoData) (string, error)
}

type koyoDataUsecaseImpl struct {
	koyoDataRepository entrepo.KoyoDataRepository
}

func NewKoyoDataUsecase(koyoDataRepository entrepo.KoyoDataRepository) KoyoDataUsecase {
	return &koyoDataUsecaseImpl{
		koyoDataRepository: koyoDataRepository,
	}
}

func (u *koyoDataUsecaseImpl) CreateKoyoData(ctx context.Context, koyoData *v1.KoyoData) (string, error) {
	if koyoData.KoyoId == nil || koyoData.KoyoScale == nil || koyoData.KoyoDataParams == nil {
		return "", ErrorPropertyNotSet
	}

	values := make([]domain.Value, len(koyoData.Content))
	for i, content := range koyoData.Content {
		point := domain.Point{Lat: content.Point.Lat, Lng: content.Point.Lng}
		values[i] = domain.Value{Point: point, Value: content.Value, Optional: content.Optional}
	}

	id, err := u.koyoDataRepository.CreateKoyoData(
		ctx,
		pulid.ID(koyoData.KoyoId.Value),
		float64(*koyoData.KoyoScale),
		koyoData.KoyoDataParams,
		values,
	)
	if err != nil {
		return "", err
	}

	return string(*id), nil
}
