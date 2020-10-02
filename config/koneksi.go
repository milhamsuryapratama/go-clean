package config

import (
	"go-clean/domain"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connect ...
func Connect() *gorm.DB {
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

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
