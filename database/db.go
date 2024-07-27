package database

import (
	"fmt"
	"os"

	"github.com/accalina/simple-loan-engine/models"
	"github.com/accalina/simple-loan-engine/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDBPostgre() {
	var err error
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
		os.Getenv("DB_TIMEZONE"),
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	utils.PanicLogging(err)

	DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	DB.AutoMigrate(
		&models.Loan{},
		&models.Approval{},
		&models.Investment{},
		&models.Investor{},
		&models.Disbursement{},
	)
}
