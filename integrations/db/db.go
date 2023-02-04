package db

import (
	"github.com/restuhaqza/tdd-in-go/integrations/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=passw0rd DB.name=test_db port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	DB = db

	return db, nil
}

func GetDB() *gorm.DB {
	return DB
}

func Migrate() {
	if DB != nil {
		DB.AutoMigrate(&entity.Note{})
	}
}

func Clean() {
	if DB != nil {
		DB.Migrator().DropTable(&entity.Note{})
	}
}
