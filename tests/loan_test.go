package tests

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/accalina/simple-loan-engine/database"
	"github.com/accalina/simple-loan-engine/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateLoan(t *testing.T) {
	app := setupTestApp()

	payload := `{
        "borrower_id": "123456789",
        "principal_amount": 10000,
        "rate": 5,
        "roi": 10,
        "agreement_letter_link": "http://example.com/agreement.pdf"
    }`

	req := httptest.NewRequest("POST", "/loan", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 201, resp.StatusCode)

	var loan models.Loan
	json.NewDecoder(resp.Body).Decode(&loan)
	assert.Equal(t, "123456789", loan.BorrowerID)
	assert.Equal(t, float64(10000), loan.PrincipalAmount)
}

func TestApproveLoan(t *testing.T) {
	app := setupTestApp()

	// Create a loan
	loanID := uuid.New()
	loan := models.Loan{
		ID:                  loanID,
		BorrowerID:          "123456789",
		PrincipalAmount:     10000,
		Rate:                5,
		ROI:                 10,
		AgreementLetterLink: "http://example.com/agreement.pdf",
		State:               "proposed",
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
	}
	database.DB.Create(&loan)

	payload := `{
        "proof_photo_url": "http://example.com/proof.jpg",
        "validator_id": "987654321",
        "approval_date": "2024-07-26T00:00:00Z"
    }`

	req := httptest.NewRequest("PUT", "/loan/approve/"+loanID.String(), bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	database.DB.First(&loan, "id = ?", loanID)
	assert.Equal(t, "approved", loan.State)
}
