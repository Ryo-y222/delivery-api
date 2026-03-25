package model

import "time"

type Qualification struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"not null"` // 例: 大型免許、危険物取扱者
	Description string    `json:"description"`          // 資格の説明
	CreatedAt   time.Time `json:"created_at"`
}
