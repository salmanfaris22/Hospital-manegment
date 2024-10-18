package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"main.go/model"
)

func DbInit() *gorm.DB {
	dsn := "host=localhost user=postgres password=poomon dbname=hospiteldb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect to the database", err)
		return nil
	}
	err = db.AutoMigrate(&model.Patient{}, &model.Appointment{}, &model.Doctor{}, &model.Medicine{})
	if err != nil {
		fmt.Println("AutoMigrate Error", err)
		return nil
	}
	return db
}
