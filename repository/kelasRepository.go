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
	m.Conn.Create(&k)
	return k, nil
}

// Show ...
func (m *KelasRepository) Show(id string) (res domain.Kelas, err error) {
	var kelas domain.Kelas
	m.Conn.First(&kelas, id)
	return kelas, nil
}

// Delete ...
func (m *KelasRepository) Delete(id string) (res domain.Kelas, err error) {
	var kelas domain.Kelas
	m.Conn.Delete(&kelas, id)
	return res, nil
}

// Update ...
func (m *KelasRepository) Update(k domain.Kelas, id string) (res domain.Kelas, err error) {
	var kelas domain.Kelas
	m.Conn.First(&kelas, id)
	kelas.NamaKelas = k.NamaKelas
	m.Conn.Save(&kelas)
	return kelas, nil
}
