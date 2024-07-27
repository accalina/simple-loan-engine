package models

import (
	"time"

	"github.com/google/uuid"
)

type Loan struct {
	ID                  uuid.UUID    `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	BorrowerID          string       `json:"borrower_id"`
	PrincipalAmount     float64      `json:"principal_amount"`
	Rate                float64      `json:"rate"`
	ROI                 float64      `json:"roi"`
	AgreementLetterLink string       `json:"agreement_letter_link"`
	State               string       `json:"state"`
	CreatedAt           time.Time    `json:"created_at"`
	UpdatedAt           time.Time    `json:"updated_at"`
	ApprovalInfo        Approval     `json:"approval_info" gorm:"foreignKey:LoanID"`
	Investments         []Investment `json:"investments" gorm:"foreignKey:LoanID"`
	DisbursementInfo    Disbursement `json:"disbursement_info" gorm:"foreignKey:LoanID"`
}

type Approval struct {
	ID            uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	LoanID        uuid.UUID `json:"loan_id" gorm:"type:uuid"`
	ProofPhotoURL string    `json:"proof_photo_url"`
	ValidatorID   string    `json:"validator_id"`
	ApprovalDate  time.Time `json:"approval_date"`
}

type Disbursement struct {
	ID                 uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	LoanID             uuid.UUID `json:"loan_id" gorm:"type:uuid"`
	SignedAgreementURL string    `json:"signed_agreement_url"`
	OfficerID          string    `json:"officer_id"`
	DisbursementDate   time.Time `json:"disbursement_date"`
}
