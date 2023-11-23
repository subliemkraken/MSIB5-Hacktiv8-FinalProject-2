package config

import (
	"FinalProject2/model/entity"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {

	dsn := "host=" + os.Getenv("PGHOST") + " user=" + os.Getenv("PGUSER") + " password=" + os.Getenv("PGPASSWORD") + " dbname=" + os.Getenv("PGDATABASE") + " port=" + os.Getenv("PGPORT") + " sslmode=disable TimeZone=Asia/Jakarta"
	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.User{})

	return db
}
