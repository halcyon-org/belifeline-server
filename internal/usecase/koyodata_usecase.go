package usecase

import (
	"context"

	"github.com/halcyon-org/kizuna/ent/schema/pulid"
	v1 "github.com/halcyon-org/kizuna/gen/belifeline/models/v1"
	entrepo "github.com/halcyon-org/kizuna/internal/adapter/repository/ent"
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

	id, err := u.koyoDataRepository.CreateKoyoData(ctx, pulid.ID(koyoData.KoyoId.Value), float64(*koyoData.KoyoScale), koyoData.KoyoDataParams)
	if err != nil {
		return "", err
	}

	return string(*id), nil
}
