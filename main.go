package main

import (
	"github.com/accalina/simple-loan-engine/database"
	"github.com/accalina/simple-loan-engine/routers"
	"github.com/accalina/simple-loan-engine/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	utils.PanicLogging(err)

	database.ConnectDBPostgre()

	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())

	routers.LoanRouters(app)
	routers.InvestorRouters(app)
	app.Listen(":8081")
}
