package model

import "time"

// Payment status constants
const (
	PaymentStatusPending   = "pending"
	PaymentStatusSucceeded = "succeeded"
	PaymentStatusFailed    = "failed"
	PaymentStatusRefunded  = "refunded"
)

// Payment はマッチングに対する決済記録を表す
type Payment struct {
	ID              uint       `json:"id" gorm:"primaryKey"`
	MatchID         uint       `json:"match_id" gorm:"index:idx_payments_match"`
	Match           Match      `json:"match" gorm:"foreignKey:MatchID" binding:"-"`
	PayerID         uint       `json:"payer_id"`
	Payer           User       `json:"payer" gorm:"foreignKey:PayerID" binding:"-"`
	Amount          int        `json:"amount"`
	Currency        string     `json:"currency" gorm:"type:varchar(10);default:jpy;not null"`
	StripePaymentID string     `json:"stripe_payment_id" gorm:"type:varchar(255)"`
	Status          string     `json:"status" gorm:"type:varchar(20);default:pending;not null"`
	PaidAt          *time.Time `json:"paid_at"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}

func (Payment) TableName() string {
	return "payments"
}
