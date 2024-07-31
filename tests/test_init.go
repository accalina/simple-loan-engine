package tests

import (
	"github.com/accalina/simple-loan-engine/database"
	"github.com/accalina/simple-loan-engine/routers"
	"github.com/accalina/simple-loan-engine/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func setupTestApp() *fiber.App {
	err := godotenv.Load("../.env")
	utils.PanicLogging(err)

	database.ConnectDBPostgre()
	app := fiber.New()
	routers.InvestorRouters(app)
	return app
}
