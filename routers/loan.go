package routers

import (
	"github.com/accalina/simple-loan-engine/handlers"
	"github.com/gofiber/fiber/v2"
)

func LoanRouters(app *fiber.App) {
	loan := app.Group("/loan")

	loan.Get("/:id", handlers.GetLoanByID)
	loan.Post("/", handlers.CreateLoan)
	loan.Put("/approve/:id", handlers.ApproveLoan)
	loan.Put("/invest/:id", handlers.InvestInLoan)
	loan.Put("/disburse/:id", handlers.DisburseLoan)
}
