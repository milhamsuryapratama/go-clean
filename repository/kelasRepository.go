package repository

import (
	"context"
	"go-clean/domain"

	"gorm.io/gorm"
)

// KelasRepository ...
type KelasRepository struct {
	Conn *gorm.DB
}

// NewKelasRepository ...
func NewKelasRepository(Conn *gorm.DB) domain.KelasRepository {
	return &KelasRepository{Conn}
}

// Get ...
func (m *KelasRepository) Get(ctx context.Context) (res []domain.Kelas, err error) {
	var kelas []domain.Kelas
	m.Conn.Find(&kelas)
	return kelas, nil
}

// Create ...
func (m *KelasRepository) Create(k domain.Kelas) (kelas domain.Kelas, err error) {
	b := m.Conn.Create(k)
	if b != nil {
		j := domain.Kelas{}
		return j, err
	}
	return k, nil
}
