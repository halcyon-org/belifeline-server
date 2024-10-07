package ent

import "github.com/halcyon-org/kizuna/ent"

type KoyoInformationRepository interface {
}

type koyoInformationRepositoryImpl struct {
	DB *ent.Client
}

func NewKoyoInformationRepository(db *ent.Client) KoyoInformationRepository {
	return &koyoInformationRepositoryImpl{
		DB: db,
	}
}
