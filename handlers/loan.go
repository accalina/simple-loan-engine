package handlers

import (
	"fmt"
	"time"

	"github.com/accalina/simple-loan-engine/database"
	"github.com/accalina/simple-loan-engine/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// for creating new Loan
func CreateLoan(c *fiber.Ctx) error {
	loan := new(models.Loan)

	err := c.BodyParser(loan)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Cannot parse JSON",
			"error":   err.Error(),
			"success": false,
		})
	}

	loan.ID = uuid.New()
	loan.State = "proposed"
	loan.CreatedAt = time.Now()
	loan.UpdatedAt = time.Now()
	database.DB.Create(&loan)

	return c.Status(201).JSON(loan)
}

func GetLoanByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Cannot parse UUID",
			"success": false,
		})
	}

	loan := new(models.Loan)
	err = database.DB.First(&loan, id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Loan not found",
			"success": false,
		})
	}
	return c.JSON(&loan)
}

// for approving existing loan
func ApproveLoan(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Cannot parse UUID",
			"success": false,
		})
	}

	// Check if loan is present
	loan := new(models.Loan)
	err = database.DB.First(&loan, id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Loan not found",
			"success": false,
		})
	}

	if loan.State != "proposed" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Loan cannot be approved in its current state",
			"success": false,
		})
	}

	approval := new(models.Approval)
	err = c.BodyParser(approval)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Cannot parse JSON",
			"success": false,
		})
	}

	approval.ID = uuid.New()
	approval.LoanID = id
	approval.ApprovalDate = time.Now()
	database.DB.Create(&approval)

	loan.State = "approved"
	loan.ApprovalInfo = *approval
	loan.UpdatedAt = time.Now()
	database.DB.Save(&loan)

	return c.JSON(loan)
}

func InvestInLoan(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Cannot parse UUID",
			"success": false,
		})
	}

	loan := new(models.Loan)
	err = database.DB.Preload("ApprovalInfo").First(&loan, id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Loan not found",
			"success": false,
		})
	}

	if loan.State != "approved" {
		return c.Status(404).JSON(fiber.Map{
			"message": "Loan is not approved",
			"success": false,
		})
	}

	investment := new(models.Investment)
	err = c.BodyParser(investment)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Cannot parse JSON",
			"detail":  err.Error(),
			"success": false,
		})
	}

	investment.ID = uuid.New()
	investment.LoanID = id
	investment.CreatedAt = time.Now()
	investment.UpdatedAt = time.Now()

	// Calculate the total invested amount
	var totalInvested float64
	database.DB.Model(&models.Investment{}).Where("loan_id = ?", id).Select("COALESCE(SUM(amount), 0)").Scan(&totalInvested)

	remainingInvest := loan.PrincipalAmount - totalInvested
	if totalInvested+investment.Amount > loan.PrincipalAmount {
		return c.Status(400).JSON(fiber.Map{
			"message": fmt.Sprintf("Total invested amount cannot exceed the principal amount, maximum amount is %.2f", remainingInvest),
			"success": false,
		})
	}
	database.DB.Create(&investment)
	loan.Investments = []models.Investment{*investment}

	// Update loan state if fully invested
	if totalInvested+investment.Amount == loan.PrincipalAmount {
		loan.State = "invested"
		loan.UpdatedAt = time.Now()
		database.DB.Save(&loan)

		// Send emails to all investors (dummy implementation)
		// In a real application, use an email service
		var investors []models.Investor
		database.DB.Model(&models.Investor{}).Find(&investors)
		for _, investor := range investors {
			fmt.Printf("Sending email to %s with agreement link %s\n", investor.Email, loan.AgreementLetterLink)
		}
	}

	return c.Status(200).JSON(loan)
}

func DisburseLoan(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Cannot parse UUID",
			"success": false,
		})
	}

	var loan models.Loan
	if err := database.DB.First(&loan, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Loan not found",
			"success": false,
		})
	}

	if loan.State != "invested" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Loan cannot be disbursed in its current state",
			"success": false,
		})
	}

	disbursement := new(models.Disbursement)
	if err := c.BodyParser(disbursement); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Cannot parse JSON",
			"success": false,
		})
	}

	disbursement.ID = uuid.New()
	disbursement.LoanID = id
	disbursement.DisbursementDate = time.Now()
	database.DB.Create(&disbursement)

	loan.State = "disbursed"
	loan.UpdatedAt = time.Now()
	database.DB.Save(&loan)
	return c.Status(200).JSON(disbursement)
}
