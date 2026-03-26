package model

import "time"

// UserViolation はユーザーの違反記録を表す
type UserViolation struct {
	ID            uint        `json:"id" gorm:"primaryKey"`
	UserID        uint        `json:"user_id" gorm:"index:idx_user_violations_user"`
	User          User        `json:"user" gorm:"foreignKey:UserID" binding:"-"`
	ChatMessageID uint        `json:"chat_message_id"`
	ChatMessage   ChatMessage `json:"chat_message" gorm:"foreignKey:ChatMessageID" binding:"-"`
	ViolationType string      `json:"violation_type" gorm:"type:varchar(30); not null"`
	Detail        string      `json:"detail"`
	CreatedAt     time.Time   `json:"created_at"`
}

// TableName specifies the table name for GORM
func (UserViolation) TableName() string {
	return "user_violations"
}
