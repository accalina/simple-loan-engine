package handlers

import (
	"time"

	"github.com/accalina/simple-loan-engine/database"
	"github.com/accalina/simple-loan-engine/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// CreateInvestor creates a new investor
func CreateInvestor(c *fiber.Ctx) error {
	investor := new(models.Investor)
	err := c.BodyParser(investor)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Cannot parse JSON",
			"success": false,
		})
	}
	investor.ID = uuid.New()
	investor.CreatedAt = time.Now()
	investor.UpdatedAt = time.Now()
	database.DB.Create(&investor)
	return c.Status(fiber.StatusCreated).JSON(investor)
}

// GetInvestorByID fetches an investor by their UUID
func GetInvestorByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Cannot parse UUID",
			"success": false,
		})
	}

	var investor models.Investor
	err = database.DB.First(&investor, id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Investor not found",
			"success": false,
		})
	}

	return c.JSON(investor)
}
