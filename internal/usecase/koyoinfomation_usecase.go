package usecase

import (
	"context"

	"github.com/halcyon-org/kizuna/ent/schema/pulid"
	v1 "github.com/halcyon-org/kizuna/gen/belifeline/models/v1"
	"github.com/halcyon-org/kizuna/gen/ent/koyoinformation"
	entrepo "github.com/halcyon-org/kizuna/internal/adapter/repository/ent"
	"github.com/halcyon-org/kizuna/internal/domain/domain"
	"github.com/halcyon-org/kizuna/internal/util"
)

type KoyoInformationUsecase interface {
	CreateKoyoInformation(ctx context.Context, koyoInformation *v1.KoyoInformation) (*domain.KoyoInformation, error)
}

type koyoInformationUsecaseImpl struct {
	koyoInformationRepository entrepo.KoyoInformationRepository
}

func NewKoyoInformationUsecase(koyoInformationRepository entrepo.KoyoInformationRepository) KoyoInformationUsecase {
	return &koyoInformationUsecaseImpl{
		koyoInformationRepository: koyoInformationRepository,
	}
}

func (u *koyoInformationUsecaseImpl) CreateKoyoInformation(ctx context.Context, koyoInformation *v1.KoyoInformation) (*domain.KoyoInformation, error) {
	if koyoInformation.KoyoName == nil || koyoInformation.KoyoDescription == nil || koyoInformation.KoyoScales == nil || koyoInformation.License == nil {
		return nil, ErrorPropertyNotSet
	}

	externalIds := make([]pulid.ID, len(koyoInformation.NeedExternal))
	for i, externalId := range koyoInformation.NeedExternal {
		externalIds[i] = pulid.ID(externalId.Value)
	}
	scales := make([]float64, len(koyoInformation.KoyoScales))
	for i, scale := range koyoInformation.KoyoScales {
		scales[i] = float64(scale)
	}
	dataIds := make([]pulid.ID, len(koyoInformation.KoyoDataIds))
	for i, dataId := range koyoInformation.KoyoDataIds {
		dataIds[i] = pulid.ID(dataId.Value)
	}

	data, err := u.koyoInformationRepository.CreateKoyoInformation(ctx, *koyoInformation.KoyoName, *koyoInformation.KoyoDescription, externalIds, koyoInformation.KoyoParams, scales, dataIds, koyoInformation.Version.Value, *koyoInformation.License, koyoinformation.DataType(koyoInformation.DataType.String()), util.GenApiKey())
	if err != nil {
		return nil, err
	}

	domainData := domain.ToDomainKoyoInformation(*data)
	return &domainData, nil
}
