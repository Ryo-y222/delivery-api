package model

import "time"

type User struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	CompanyID    *uint     `json:"company_id"`
	Company      Company   `json:"company" gorm:"foreignKey:CompanyID" binding:"-"`
	Email        string    `json:"email" gorm:"type:varchar(255);uniqueIndex;not null" binding:"required"`
	PasswordHash string    `json:"-" gorm:"not null"`
	Name         string    `json:"name" gorm:"not null" binding:"required"`
	Role         string    `json:"role" gorm:"type:varchar(20);default:shipper;not null"`
	Phone        string    `json:"phone" gorm:"type:varchar(20)"`
	IsActive     bool      `json:"is_active" gorm:"default:true"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (User) TableName() string { return "users" }
