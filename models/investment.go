package models

import (
	"time"

	"github.com/google/uuid"
)

type Investment struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	LoanID     uuid.UUID `json:"loan_id" gorm:"type:uuid"`
	InvestorID uuid.UUID `json:"investor_id" gorm:"type:uuid"`
	Amount     float64   `json:"amount"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Investor struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
