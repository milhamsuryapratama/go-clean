package test

import (
	"database/sql"
	"go-clean/domain"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// KelasT ...
type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository domain.KelasRepository
	kelas      *domain.Kelas
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("mysql", db)
	require.NoError(s.T(), err)

	s.DB.LogMode(true)

	// s.repository = CreateRepository(s.DB)
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

// TestCreateKelas ...
// func (s *Suite) TestCreateKelas() {
// 	k := s.kelas{
// 		NamaKelas: "IPA",
// 	}

// 	p, _ := s.repository.Create(k)

// 	assert.Equal(s.T(), "IPA", p.NamaKelas)
// }
