package domain

import (
	"context"
)

// Siswa ...
type Siswa struct {
	ID      int `gorm:"primary_key;auto_increment"`
	Nama    string
	Jk      string
	Alamat  string
	KelasID int   `sql:"index"`
	Kelas   Kelas `gorm:"foreignkey:KelasID"`
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
}

// SiswaRepository ...
type SiswaRepository interface {
	Get(ctx context.Context) (res []Siswa, err error)
	Show(id string) (Siswa, error)
	Create(siswa Siswa) (Siswa, error)
}
