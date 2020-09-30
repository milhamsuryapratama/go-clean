package repository

import (
	"context"
	"go-clean/domain"

	"gorm.io/gorm"
)

// SiswaRepository ...
type SiswaRepository struct {
	Conn *gorm.DB
}

// NewSiswaRepository ...
func NewSiswaRepository(Conn *gorm.DB) domain.SiswaRepository {
	return &SiswaRepository{Conn}
}

// Get ...
func (s *SiswaRepository) Get(ctx context.Context) (res []domain.Siswa, err error) {
	var siswa []domain.Siswa
	s.Conn.Preload("Kelas").Find(&siswa)
	return siswa, nil
}

// Show ...
func (s *SiswaRepository) Show(id string) (res domain.Siswa, err error) {
	var siswa domain.Siswa
	s.Conn.Preload("Kelas").First(&siswa, id)
	return siswa, nil
}
