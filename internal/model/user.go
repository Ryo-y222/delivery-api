package model

import "time"

type User struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Email        string    `json:"email" gorm:"type:varchar(255);uniqueIndex" binding:"required"`
	PasswordHash string    `json:"-"`
	Name         string    `json:"name" binding:"required"`
	Role         string    `json:"role" gorm:"default:shipper"`
	Company      string    `json:"company"`
	Phone        string    `json:"phone" gorm:"type:varchar(20)"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (User) TableName() string { return "users" }
