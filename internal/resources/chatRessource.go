package resources

import "time"

type ChatRessource struct {
	Success bool `json:"success,omitempty"`
	Data    struct {
		Sender_id         int        `json:"sender_id"`
		Receiver_id       int        `json:"receiver_id"`
		Receiver_username string     `json:"receiver_username"`
		CreatedAt         *time.Time `json:"created_at"`
		Id                int        `json:"id"`
		Message           string     `json:"message"`
	} `json:"data"`
}
