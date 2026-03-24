package model

import "time"

// DispatchPlan status constants
const (
	DispatchPlanStatusPlanned    = "planned"
	DispatchPlanStatusInProgress = "in_progress"
	DispatchPlanStatusCompleted  = "completed"
	DispatchPlanStatusCancelled  = "cancelled"
)

// DispatchPlan は「1台のトラック × 1日」の配車計画を表す
type DispatchPlan struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CompanyID uint      `json:"company_id"`
	Company   Company   `json:"company" gorm:"foreignKey:CompanyID" binding:"-"`
	VehicleID uint      `json:"vehicle_id"`
	Vehicle   Vehicle   `json:"vehicle" gorm:"foreignKey:VehicleID" binding:"-"`
	DriverID  uint      `json:"driver_id"`
	Driver    User      `json:"driver" gorm:"foreignKey:DriverID" binding:"-"`
	PlanDate  time.Time `json:"plan_date" gorm:"type:date;index:idx_dispatch_plans_company_date"`
	Status    string    `json:"status" gorm:"type:varchar(20);default:planned;index:idx_dispatch_plans_status"`
	Note      string    `json:"note"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (DispatchPlan) TableName() string {
	return "dispatch_plans"
}
