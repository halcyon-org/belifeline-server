package ent

import "github.com/halcyon-org/kizuna/ent"

type KoyoInfomationRepository interface {
}

type koyoInfomationRepositoryImpl struct {
	DB *ent.Client
}

func NewKoyoInfomationRepository(db *ent.Client) KoyoInfomationRepository {
	return &koyoInfomationRepositoryImpl{
		DB: db,
	}
}
