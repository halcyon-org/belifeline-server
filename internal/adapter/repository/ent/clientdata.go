package ent

import "github.com/halcyon-org/kizuna/ent"

type ClientDataRepository interface {
}

type clientDataRepositoryImpl struct {
	DB *ent.Client
}

func NewClientDataRepository(db *ent.Client) ClientDataRepository {
	return &clientDataRepositoryImpl{
		DB: db,
	}
}
