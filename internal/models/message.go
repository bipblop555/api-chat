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
	Sender    int        `gorm:"not_null" validate:"required"`
	Receiver  int        `gorm:"not_null" validate:"required"`
	Message   string     `gorm:"not_null" validate:"required"`
	CreatedAt *time.Time `gorm:"type:timestamp"`
	UpdatedAt *time.Time `gorm:"type:timestamp;autoUpdateTime:true"`
}

type Sender struct {
	Sender int `gorm:"not_null" validate:"required"`
}

func (ug *DbGorm) CreateMessage(entity interface{}, w http.ResponseWriter) error {
	db := ug.Db.Table("messages").Create(entity)

	if db.Error != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return db.Error
	}

	w.WriteHeader(http.StatusCreated)
	return nil
}

func (ug *DbGorm) GetAllLinkedChat(senderID int) ([]Message, error) {
	var messages []Message

	var result []struct {
		Receiver string
		Username string
		Message  string
	}

	db := ug.Db.Model(&Message{}).
		Select("messages.receiver, users.username, messages.message").
		Joins("JOIN users ON messages.receiver = users.id").
		Where("messages.sender = ? AND messages.deleted_at IS NULL AND messages.created_at = (SELECT MAX(created_at) FROM messages WHERE sender = ? AND receiver = messages.receiver AND deleted_at IS NULL LIMIT 50)", senderID, senderID).
		Find(&messages)

	if db.Error != nil {
		return nil, nil
	}

	// Affichez les résultats.
	for _, row := range result {
		fmt.Printf("Receiver: %s, Username: %s, Message: %s\n", row.Receiver, row.Username, row.Message)
	}

	return messages, nil
}
