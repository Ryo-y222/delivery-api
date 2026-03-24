package model

import "time"

// Match status constants
const (
	MatchStatusPending        = "pending"
	MatchStatusApproved       = "approved"
	MatchStatusPaymentPending = "payment_pending"
	MatchStatusCompleted      = "completed"
	MatchStatusRejected       = "rejected"
	MatchStatusCancelled      = "cancelled"
)

// Match request type constants
const (
	RequestTypeShipperToCompany = "shipper_to_company"
	RequestTypeCompanyToCompany = "company_to_company"
)

// Match は荷主/運送会社からの区間マッチングリクエストを表す
type Match struct {
	ID                 uint      `json:"id" gorm:"primaryKey"`
	TripLegID          uint      `json:"trip_leg_id" gorm:"index:idx_matches_trip_leg"`
	TripLeg            TripLeg   `json:"trip_leg" gorm:"foreignKey:TripLegID" binding:"-"`
	RequesterID        uint      `json:"requester_id" gorm:"index:idx_matches_requester"`
	Requester          User      `json:"requester" gorm:"foreignKey:RequesterID" binding:"-"`
	RequesterCompanyID *uint     `json:"requester_company_id"`
	RequesterCompany   Company   `json:"requester_company" gorm:"foreignKey:RequesterCompanyID" binding:"-"`
	CargoWeight        float64   `json:"cargo_weight"`
	CargoDescription   string    `json:"cargo_description"`
	Message            string    `json:"message"`
	Status             string    `json:"status" gorm:"type:varchar(20);default:pending;index:idx_matches_status"`
	RejectReason       string    `json:"reject_reason"`
	RequestType        string    `json:"request_type" gorm:"type:varchar(30);default:shipper_to_company"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

func (Match) TableName() string {
	return "matches"
}
