package model

import "time"

type VehicleInspection struct {
	ID             uint       `json:"id" gorm:"primaryKey"`
	VehicleID      uint       `json:"vehicle_id" gorm:"index"`
	Vehicle        Vehicle    `json:"vehicle" gorm:"foreignKey:VehicleID" binding:"-"`
	InspectionType string     `json:"inspection_type" gorm:"type:varchar(20);not null"` // vehicle_inspection(車検) / quarterly(3ヶ月点検) / daily(日常点検)
	ScheduledAt    time.Time  `json:"scheduled_at"`                                     // 予定日
	CompletedAt    *time.Time `json:"completed_at"`                                     // 完了日（未完了はNULL）
	IsReserved     bool       `json:"is_reserved" gorm:"default:false"`                 // 予約済みフラグ
	ReservedAt     *time.Time `json:"reserved_at"`                                      // 予約日時
	ShopName       string     `json:"shop_name"`                                        // 整備工場名
	Cost           int        `json:"cost"`                                             // 費用
	Note           string     `json:"note"`                                             // 備考
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}

func (VehicleInspection) TableName() string {
	return "vehicle_inspections"
}
