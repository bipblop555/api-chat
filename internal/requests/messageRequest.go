package requests

type SendMessageRequest struct {
	Success bool `json:"success,omitempty"`
	Data    struct {
		Type       string `json:"type"`
		Attributes struct {
			Sender_id   int    `json:"sender_id" validate:"required"`
			Receiver_id int    `json:"receiver_id" validate:"required"`
			Message     string `json:"message" validate:"required"`
		} `json:"attributes"`
	} `json:"data"`
}
