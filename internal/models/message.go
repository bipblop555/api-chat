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

	fmt.Print(result)

	query := `
        SELECT messages.id, messages.message, messages.sender, messages.receiver, messages.created_at, messages.deleted_at, users.username
        FROM messages
        JOIN users ON messages.receiver = users.id
        JOIN (
            SELECT id, ROW_NUMBER() OVER (PARTITION BY receiver ORDER BY created_at DESC) AS rnum
            FROM messages
            WHERE deleted_at IS NULL
        ) AS ranked ON ranked.id = messages.id
        WHERE ranked.rnum = 1 AND messages.sender = ?
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
