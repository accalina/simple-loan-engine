package routers

import (
	"github.com/accalina/simple-loan-engine/handlers"
	"github.com/gofiber/fiber/v2"
)

func InvestorRouters(app *fiber.App) {
	investor := app.Group("/investor")

	investor.Post("/", handlers.CreateInvestor)
	investor.Get("/:id", handlers.GetInvestorByID)
}
