package config

import (
	"go-clean/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect ...
func Connect() *gorm.DB {
	// dsn := "root:@tcp(127.0.0.1:3306)/go-api?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "user=postgres password=ilham21 dbname=go-api port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	hasKelasTable := db.Migrator().HasTable(domain.Kelas{})
	hasSiswaTable := db.Migrator().HasTable(domain.Siswa{})
	if !hasKelasTable && !hasSiswaTable {
		db.AutoMigrate(domain.Kelas{})
		db.AutoMigrate(domain.Siswa{})
	}
	// db.Migrator().DropTable(domain.Kelas{})
	// db.Migrator().DropTable(domain.Siswa{})
	return db
}
