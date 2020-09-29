package controller

import (
	"context"
	"go-clean/domain"
)

// KelasEntity ...
type KelasEntity struct {
	kelasRepo domain.KelasRepository
}

// NewKelasEntity ...
func NewKelasEntity(a domain.KelasRepository) domain.KelasEntity {
	return &KelasEntity{
		kelasRepo: a,
	}
}

// Get ...
func (k *KelasEntity) Get(c context.Context) (res []domain.Kelas, err error) {
	res, err = k.kelasRepo.Get(c)
	if err != nil {
		return nil, err
	}
	return
}

// Create ...
func (k *KelasEntity) Create(c domain.Kelas) (kelas domain.Kelas, err error) {
	kelas, err = k.kelasRepo.Create(c)
	if err != nil {
		return kelas, err
	}

	return
}