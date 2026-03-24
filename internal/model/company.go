package model

import "time"

type Company struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	Name            string    `json:"name" gorm:"type:varchar(255)" binding:"required"`
	DisplayName     string    `json:"display_name" gorm:"type:varchar(255)"`
	Address         string    `json:"address"`
	HeadquartersLat float64   `json:"headquarters_lat"`
	HeadquartersLng float64   `json:"headquarters_lng"`
	Phone           string    `json:"phone" gorm:"type:varchar(20)"`
	Email           string    `json:"email" gorm:"type:varchar(255)"`
	LogoURL         string    `json:"logo_url"`
	RatingAvg       float64   `json:"rating_avg" gorm:"default:0"`
	TotalDeals      int       `json:"total_deals" gorm:"default:0"`
	Plan            string    `json:"plan" gorm:"type:varchar(20);default:free"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (Company) TableName() string {
	return "companies"
}
