package domain

import (
	"context"

	"gorm.io/gorm"
)

// Siswa ...
type Siswa struct {
	gorm.Model
	Nama    string
	Jk      string
	Alamat  string
	KelasID int   `sql:"index"`
	Kelas   Kelas `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// Tabler ...
type Tabler interface {
	TableName() string
}

// TableName ...
func (Siswa) TableName() string {
	return "siswa"
}

// SiswaEntity ...
type SiswaEntity interface {
	Get(ctx context.Context) ([]Siswa, error)
	Show(id string) (Siswa, error)
	Create(siswa Siswa) (Siswa, error)
	Update(siswa Siswa, id string) (Siswa, error)
	Delete(id string) (Siswa, error)
}

// SiswaRepository ...
type SiswaRepository interface {
	Get(ctx context.Context) (res []Siswa, err error)
	Show(id string) (Siswa, error)
	Create(siswa Siswa) (Siswa, error)
	Update(siswa Siswa, id string) (Siswa, error)
	Delete(id string) (Siswa, error)
}
