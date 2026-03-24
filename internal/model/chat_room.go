package model

import "time"

// ChatRoom はマッチングに紐づくチャットルームを表す（1マッチング = 1ルーム）
type ChatRoom struct {
	ID           uint          `json:"id" gorm:"primaryKey"`
	MatchID      uint          `json:"match_id" gorm:"uniqueIndex"`
	Match        Match         `json:"match" gorm:"foreignKey:MatchID" binding:"-"`
	Status       string        `json:"status" gorm:"type:varchar(20);default:active"`
	ChatMessages []ChatMessage `json:"chat_messages" gorm:"foreignKey:ChatRoomID" binding:"-"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
}

func (ChatRoom) TableName() string {
	return "chat_rooms"
}
