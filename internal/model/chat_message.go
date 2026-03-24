package model

import "time"

// ChatMessage status constants
const (
	ChatMessageStatusSent     = "sent"     // 正常に送信された
	ChatMessageStatusFiltered = "filtered" // フィルタリングでブロックされた
	ChatMessageStatusWarned   = "warned"   // 警告付きで送信された
)

// ChatMessage はチャットルーム内の1つのメッセージを表す
type ChatMessage struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	ChatRoomID     uint      `json:"chat_room_id" gorm:"index:idx_chat_messages_room_created"`
	ChatRoom       ChatRoom  `json:"chat_room" gorm:"foreignKey:ChatRoomID" binding:"-"`
	SenderID       uint      `json:"sender_id"`
	Sender         User      `json:"sender" gorm:"foreignKey:SenderID" binding:"-"`
	Content        string    `json:"content" gorm:"type:text"`
	Status         string    `json:"status" gorm:"type:varchar(20);default:sent"`
	FilteredReason string    `json:"filtered_reason"`
	CreatedAt      time.Time `json:"created_at" gorm:"index:idx_chat_messages_room_created"`
}

func (ChatMessage) TableName() string {
	return "chat_messages"
}
