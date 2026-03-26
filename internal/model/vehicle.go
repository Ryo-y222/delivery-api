package model

import "time"

// Vehicle は運送会社が所有する車両を表す
type Vehicle struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	CompanyID   uint      `json:"company_id"`
	Company     Company   `json:"company" gorm:"foreignKey:CompanyID" binding:"-"`
	PlateNumber string    `json:"plate_number" gorm:"type:varchar(20);uniqueIndex;not null" binding:"required"`
	VehicleType string    `json:"vehicle_type" gorm:"type:varchar(20);not null" binding:"required"`
	MaxWeight   float64   `json:"max_weight"`
	Status      string    `json:"status" gorm:"type:varchar(20);default:active;not null"`
	Note        string    `json:"note"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Vehicle) TableName() string {
	return "vehicles"
}
