package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connect ...
func Connect() *gorm.DB {
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db
}
