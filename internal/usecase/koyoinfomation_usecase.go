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

func (u *koyoInformationUsecaseImpl) CreateKoyoInformation(
	ctx context.Context,
	koyoInformation *v1.KoyoInformation,
) (*domain.KoyoInformation, error) {
	if koyoInformation.KoyoName == nil ||
		koyoInformation.KoyoDescription == nil ||
		koyoInformation.KoyoScales == nil ||
		koyoInformation.License == nil {
		return nil, ErrPropertyNotSet
	}

	externalIds := make([]pulid.ID, len(koyoInformation.GetNeedExternal()))
	for i, externalId := range koyoInformation.GetNeedExternal() {
		externalIds[i] = pulid.ID(externalId.GetValue())
	}

	scales := make([]float64, len(koyoInformation.GetKoyoScales()))
	for i, scale := range koyoInformation.GetKoyoScales() {
		scales[i] = float64(scale)
	}

	dataIds := make([]pulid.ID, len(koyoInformation.GetKoyoDataIds()))
	for i, dataId := range koyoInformation.GetKoyoDataIds() {
		dataIds[i] = pulid.ID(dataId.GetValue())
	}

	data, err := u.koyoInformationRepository.CreateKoyoInformation(
		ctx,
		koyoInformation.GetKoyoName(),
		koyoInformation.GetKoyoDescription(),
		externalIds,
		koyoInformation.GetKoyoParams(),
		scales,
		dataIds,
		koyoInformation.GetVersion().GetValue(),
		koyoInformation.GetLicense(),
		koyoinformation.DataType(koyoInformation.GetDataType().String()),
		util.GenAPIKey(),
	)
	if err != nil {
		return nil, err
	}

	domainData := domain.ToDomainKoyoInformation(*data)

	return &domainData, nil
}
