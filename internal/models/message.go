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

func (ug *DbGorm) GetAllLinkedChat(senderID int) ([]Message, error) {
	var messages []Message

	var result []struct {
		Receiver string
		Username string
		Message  string
	}

	fmt.Print(result)

	query := `
        SELECT messages.id, messages.message, messages.sender_id, messages.receiver_id, messages.created_at, messages.deleted_at, users.username
        FROM messages
        JOIN users ON messages.receiver_id = users.id
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

	// Affichez les résultats.
	for _, row := range result {
		fmt.Printf("Receiver: %s, Username: %s, Message: %s\n", row.Receiver, row.Username, row.Message)
	}

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
