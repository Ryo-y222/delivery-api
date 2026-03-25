package model

import "time"

// BlockedPattern はチャットでブロックする正規表現パターンを表す
type BlockedPattern struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Pattern     string    `json:"pattern" gorm:"type:varchar(500);not null" binding:"required"`
	Description string    `json:"description" gorm:"type:varchar(255)"`
	IsActive    bool      `json:"is_active" gorm:"default:true"`
	CreatedAt   time.Time `json:"created_at"`
}

func (BlockedPattern) TableName() string {
	return "blocked_patterns"
}
