package domain

import (
	"context"
)

// Kelas is ...
type Kelas struct {
	ID        int
	NamaKelas string `json:"nama_kelas"`
}

// KelasEntity ...
type KelasEntity interface {
	Get(ctx context.Context) ([]Kelas, error)
	Create(kelas Kelas) (Kelas, error)
	Show(id string) (Kelas, error)
}

// KelasRepository ...
type KelasRepository interface {
	Get(ctx context.Context) (res []Kelas, err error)
	Create(kelas Kelas) (Kelas, error)
	Show(id string) (Kelas, error)
}
