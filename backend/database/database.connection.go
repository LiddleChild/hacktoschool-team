package database

import (
	"log"

	"github.com/thamph/thinc-hacktoschool/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbInstance struct {
	Db *gorm.DB
}

var DB DbInstance

func ConnectDatabase() {
	connectionString := "host=database user=mepua password=123456 dbname=e_course port=5432 sslmode=disable TimeZone=Asia/Bangkok"

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect database")
	}

	log.Print("database is connected")

	db.AutoMigrate(&models.User{}, &models.Course{}, &models.Enrollment{}, &models.Student{}, &models.Instructor{})

	DB = DbInstance{
		Db: db,
	}
}
