package model

import "time"

type DriverQualification struct {
	ID              uint          `json:"id" gorm:"primaryKey"`
	DriverProfileID uint          `json:"driver_profile_id" gorm:"index"`
	QualificationID uint          `json:"qualification_id" gorm:"index"`
	Qualification   Qualification `json:"qualification" gorm:"foreignKey:QualificationID" binding:"-"`
	AcquiredAt      time.Time     `json:"acquired_at"` // 取得日
	ExpiryAt        *time.Time    `json:"expiry_at"`   // 有効期限（期限なしはNULL）
	ImageURL        string        `json:"image_url"`   // 資格証写真URL
	CreatedAt       time.Time     `json:"created_at"`
}

func (DriverQualification) TableName() string {
	return "driver_qualifications"
}
