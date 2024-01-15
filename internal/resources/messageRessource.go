package resources

import "time"

type MessageRessource struct {
	Success bool `json:"success,omitempty"`
	Data    struct {
		Sender_id int        `json:"sender_id"`
		Message   string     `json:"message"`
		Id        int        `json:"id"`
		CreatedAt *time.Time `json:"created_at"`
	} `json:"data"`
}
