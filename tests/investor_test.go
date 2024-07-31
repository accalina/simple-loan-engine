package tests

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/accalina/simple-loan-engine/database"
	"github.com/accalina/simple-loan-engine/models"
	"github.com/stretchr/testify/assert"
)

func TestCreateInvestor(t *testing.T) {
	app := setupTestApp()

	payload := `{
        "name": "John Doe",
        "email": "johndoe@example.com"
    }`

	req := httptest.NewRequest("POST", "/investor", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 201, resp.StatusCode)

	var investor models.Investor
	json.NewDecoder(resp.Body).Decode(&investor)
	assert.Equal(t, "John Doe", investor.Name)
	assert.Equal(t, "johndoe@example.com", investor.Email)
}

func TestGetInvestorByID(t *testing.T) {
	app := setupTestApp()

	// Create an investor
	investor := models.Investor{
		Name:  "Jane Doe",
		Email: "janedoe@example.com",
	}
	database.DB.Create(&investor)

	req := httptest.NewRequest("GET", "/investor/"+investor.ID.String(), nil)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	var fetchedInvestor models.Investor
	json.NewDecoder(resp.Body).Decode(&fetchedInvestor)
	assert.Equal(t, investor.Name, fetchedInvestor.Name)
	assert.Equal(t, investor.Email, fetchedInvestor.Email)
}
