package model

import "time"

// Tracking は配車計画の位置情報記録を表す
type Tracking struct {
	ID             uint         `json:"id" gorm:"primaryKey"`
	DispatchPlanID uint         `json:"dispatch_plan_id" gorm:"index:idx_trackings_plan_recorded"`
	DispatchPlan   DispatchPlan `json:"dispatch_plan" gorm:"foreignKey:DispatchPlanID" binding:"-"`
	TripLegID      *uint        `json:"trip_leg_id"`
	TripLeg        TripLeg      `json:"trip_leg" gorm:"foreignKey:TripLegID" binding:"-"`
	Lat            float64      `json:"lat"`
	Lng            float64      `json:"lng"`
	RecordedAt     time.Time    `json:"recorded_at" gorm:"index:idx_trackings_plan_recorded"`
}

func (Tracking) TableName() string {
	return "trackings"
}
