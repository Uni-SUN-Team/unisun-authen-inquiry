package database

import (
	"log"
	"os"
	"strings"
	"unisun/api/unisun-authen-management-information/src/configs/environment"
	"unisun/api/unisun-authen-management-information/src/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dbEnv := environment.ENV.Database
	str := []string{
		"host=" + os.Getenv(dbEnv.Host),
		"user=" + os.Getenv(dbEnv.User),
		"password=" + os.Getenv(dbEnv.Pass),
		"dbname=" + os.Getenv(dbEnv.Name),
		"port=" + os.Getenv(dbEnv.Port),
		"TimeZone=" + os.Getenv(dbEnv.TimeZone),
	}
	dsn := strings.Join(str, " ")
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Failed to connect to database!")
	}

	var UserAuthPermission entities.UserAuthPermission
	database.AutoMigrate(&UserAuthPermission)
	DB = database
}
