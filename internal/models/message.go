package models

import (
	"fmt"
	"net/http"
	"time"
)

var LinkedChatResult []struct {
	Receiver string
	Username string
	Message  string
}

type Message struct {
	Sender_id   int        `gorm:"not_null" validate:"required"`
	Receiver_id int        `gorm:"not_null" validate:"required"`
	Message     string     `gorm:"not_null" validate:"required"`
	CreatedAt   *time.Time `gorm:"type:timestamp"`
	UpdatedAt   *time.Time `gorm:"type:timestamp;autoUpdateTime:true"`
}

type MessageChat struct {
	Sender_id         int        `gorm:"not_null" validate:"required"`
	Receiver_id       int        `gorm:"not_null" validate:"required"`
	Message           string     `gorm:"not_null" validate:"required"`
	Receiver_username string     `gorm:"no_null"`
	CreatedAt         *time.Time `gorm:"type:timestamp"`
	UpdatedAt         *time.Time `gorm:"type:timestamp;autoUpdateTime:true"`
}

type Sender struct {
	Sender_id int `gorm:"not_null" validate:"required"`
}

func (ug *DbGorm) CreateMessage(entity interface{}, w http.ResponseWriter) error {
	fmt.Print("CREATEMESSAGE ENTITY = ", entity)

	db := ug.Db.Table("messages").Create(entity)

	if db.Error != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return db.Error
	}

	w.WriteHeader(http.StatusCreated)
	return nil
}

func (ug *DbGorm) GetAllLinkedChat(senderID int) ([]MessageChat, error) {
	var messages []MessageChat

	query := `
    SELECT messages.id, messages.message, messages.sender_id, messages.receiver_id, messages.created_at, messages.deleted_at,
           sender.username AS sender_username, receiver.username AS receiver_username
    FROM messages
    JOIN users AS sender ON messages.sender_id = sender.id
    JOIN users AS receiver ON messages.receiver_id = receiver.id
    JOIN (
        SELECT id, ROW_NUMBER() OVER (PARTITION BY receiver_id ORDER BY created_at DESC) AS rnum
        FROM messages
        WHERE deleted_at IS NULL
    ) AS ranked ON ranked.id = messages.id
    WHERE ranked.rnum = 1 AND messages.sender_id = ?
`

	db := ug.Db.Raw(query, senderID).Scan(&messages)

	if db.Error != nil {
		return nil, nil
	}

	fmt.Print(messages)

	return messages, nil
}

func (ug *DbGorm) GetAllMessagesFromUser(senderId string, receiverId string) ([]Message, error) {
	var messages []Message
	db := ug.Db.Table("messages").Where("sender_id = ? AND receiver_id = ?", senderId, receiverId).Limit(1000).Order("created_at").Scan(&messages)
	if db.Error != nil {
		return nil, db.Error
	}

	return messages, nil
}
