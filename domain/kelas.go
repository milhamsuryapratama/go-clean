package domain

import (
	"context"
)

// Kelas is ...
type Kelas struct {
	ID        int     `gorm:"primary_key;auto_increment"`
	NamaKelas string  `json:"nama_kelas"`
	Siswa     []Siswa `gorm:"[]foreignkey:KelasID"`
}

// KelasEntity ...
type KelasEntity interface {
	Get(ctx context.Context) ([]Kelas, error)
	Create(kelas Kelas) (Kelas, error)
	Show(id string) (Kelas, error)
	Delete(id string) (Kelas, error)
	Update(kelas Kelas, id string) (Kelas, error)
}

// KelasRepository ...
type KelasRepository interface {
	Get(ctx context.Context) (res []Kelas, err error)
	Create(kelas Kelas) (Kelas, error)
	Show(id string) (Kelas, error)
	Delete(id string) (Kelas, error)
	Update(kelas Kelas, id string) (Kelas, error)
}
