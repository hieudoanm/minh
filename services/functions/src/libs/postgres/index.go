package postgres

import (
	"chatbot-functions/src/utils"
	"fmt"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

var DATABASE_HOST = utils.Getenv("DATABASE_HOST", "localhost")
var DATABASE_PORT = utils.Getenv("DATABASE_PORT", "5432")
var DATABASE_USER = utils.Getenv("DATABASE_USER", "gouser")
var DATABASE_PASS = utils.Getenv("DATABASE_PASS", "gopass")
var DATABASE_NAME = utils.Getenv("DATABASE_NAME", "postgres")
var DATABASE_MODE = utils.Getenv("DATABASE_MODE", "disable")
var DATABASE_TIMEZONE = utils.Getenv("DATABASE_TIMEZONE", "Asia/Ho_Chi_Minh")

type DatabaseConfigs struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Pass     string `json:"pass"`
	Name     string `json:"name"`
	Mode     string `json:"mode"`
	TimeZone string `json:"timeZone"`
}

var postgresDatabase *gorm.DB

func OpenDatabase(configs DatabaseConfigs) *gorm.DB {
	if postgresDatabase != nil {
		return postgresDatabase
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		configs.Host,
		configs.Port,
		configs.User,
		configs.Pass,
		configs.Name,
		configs.Mode,
		configs.TimeZone,
	)

	database, databaseError := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if databaseError != nil {
		panic(databaseError)
	}

	postgresDatabase = database
	return postgresDatabase
}

func GetDatabase() *gorm.DB {
	return OpenDatabase(DatabaseConfigs{
		Host:     DATABASE_HOST,
		Port:     DATABASE_PORT,
		User:     DATABASE_USER,
		Pass:     DATABASE_PASS,
		Name:     DATABASE_NAME,
		Mode:     DATABASE_MODE,
		TimeZone: DATABASE_TIMEZONE,
	})
}
