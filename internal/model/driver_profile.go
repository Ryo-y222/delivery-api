package model

import "time"

type DriverProfile struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	UserID          uint      `json:"user_id" gorm:"uniqueIndex"` // 1ユーザー1プロフィール
	User            User      `json:"user" gorm:"foreignKey:UserID" binding:"-"`
	LicenseNumber   string    `json:"license_number"`                  // 免許証番号
	LicenseType     string    `json:"license_type"`                    // 免許種別（大型/中型/普通等）
	LicenseExpiry   time.Time `json:"license_expiry"`                  // 免許有効期限
	LicenseImageURL string    `json:"license_image_url"`               // 免許証写真URL
	BirthDate       time.Time `json:"birth_date"`                      // 生年月日（年齢算出）
	AccidentCount   int       `json:"accident_count" gorm:"default:0"` // 事故歴件数
	AccidentDetail  string    `json:"accident_detail"`                 // 事故歴詳細
	HireDate        time.Time `json:"hire_date"`                       // 入社日
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (DriverProfile) TableName() string {
	return "driver_profiles"
}
