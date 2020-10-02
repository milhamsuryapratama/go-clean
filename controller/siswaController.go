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

// Show ...
func (s *SiswaEntity) Show(id string) (siswa domain.Siswa, err error) {
	siswa, err = s.siswaRepo.Show(id)
	if err != nil {
		return domain.Siswa{}, err
	}

	return
}

// Create ...
func (s *SiswaEntity) Create(c domain.Siswa) (siswa domain.Siswa, err error) {
	siswa, err = s.siswaRepo.Create(c)
	if err != nil {
		return siswa, err
	}

	return
}

// Update ...
func (s *SiswaEntity) Update(c domain.Siswa, id string) (siswa domain.Siswa, err error) {
	siswa, err = s.siswaRepo.Update(c, id)
	if err != nil {
		return domain.Siswa{}, err
	}

	return
}
