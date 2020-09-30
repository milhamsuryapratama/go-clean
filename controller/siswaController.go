package controller

import (
	"context"
	"go-clean/domain"
)

// SiswaEntity ...
type SiswaEntity struct {
	siswaRepo domain.SiswaRepository
}

// NewSiswaEntity ...
func NewSiswaEntity(a domain.SiswaRepository) domain.SiswaEntity {
	return &SiswaEntity{
		siswaRepo: a,
	}
}

// Get ...
func (s *SiswaEntity) Get(c context.Context) (res []domain.Siswa, err error) {
	res, err = s.siswaRepo.Get(c)
	if err != nil {
		return nil, err
	}

	return
}
