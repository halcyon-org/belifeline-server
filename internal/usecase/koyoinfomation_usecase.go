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
	UpdateKoyoInformation(ctx context.Context, koyoInformation *v1.KoyoInformation) (*domain.KoyoInformation, error)
	ListKoyoInformation(ctx context.Context, limit int32) ([]*domain.KoyoInformation, error)
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

	data, err := u.koyoInformationRepository.CreateKoyoInformation(ctx, *koyoInformation.KoyoName, *koyoInformation.KoyoDescription, externalIds, koyoInformation.KoyoParams, scales, dataIds, koyoInformation.Version.Value, *koyoInformation.License, koyoinformation.DataType(koyoInformation.DataType.String()), util.GenAPIKey())
	if err != nil {
		return nil, err
	}

	domainData := domain.ToDomainKoyoInformation(*data)
	return &domainData, nil
}

func (u *koyoInformationUsecaseImpl) UpdateKoyoInformation(ctx context.Context, koyoInformation *v1.KoyoInformation) (*domain.KoyoInformation, error) {
	var externalIds []pulid.ID
	if koyoInformation.KoyoScales != nil {
		externalIds = make([]pulid.ID, len(koyoInformation.NeedExternal))
		for i, externalId := range koyoInformation.NeedExternal {
			externalIds[i] = pulid.ID(externalId.Value)
		}
	}
	var scales []float64
	if koyoInformation.KoyoScales != nil {
		scales = make([]float64, len(koyoInformation.KoyoScales))
		for i, scale := range koyoInformation.KoyoScales {
			scales[i] = float64(scale)
		}
	}
	var dataIds []pulid.ID
	if koyoInformation.KoyoScales != nil {
		dataIds := make([]pulid.ID, len(koyoInformation.KoyoDataIds))
		for i, dataId := range koyoInformation.KoyoDataIds {
			dataIds[i] = pulid.ID(dataId.Value)
		}
	}
	var dataType koyoinformation.DataType
	if koyoInformation.DataType != nil {
		dataType = koyoinformation.DataType(koyoInformation.DataType.String())
	}

	data, err := u.koyoInformationRepository.UpdateKoyoInformation(ctx, pulid.ID(koyoInformation.KoyoId.Value), koyoInformation.KoyoName, koyoInformation.KoyoDescription, &externalIds, &koyoInformation.KoyoParams, &scales, &dataIds, &koyoInformation.Version.Value, koyoInformation.License, &dataType)
	if err != nil {
		return nil, err
	}

	domainData := domain.ToDomainKoyoInformation(*data)
	return &domainData, nil
}

func (u *koyoInformationUsecaseImpl) ListKoyoInformation(ctx context.Context, limit int32) ([]*domain.KoyoInformation, error) {
	dataList, err := u.koyoInformationRepository.GetAllKoyoInformation(ctx, limit)
	if err != nil {
		return nil, err
	}

	var domainDataList = make([]*domain.KoyoInformation, len(dataList))
	for i, data := range dataList {
		domainData := domain.ToDomainKoyoInformation(*data)
		domainDataList[i] = &domainData
	}
	return domainDataList, nil
}
