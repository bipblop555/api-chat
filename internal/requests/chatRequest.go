package requests

type ChatRequest struct {
	Success bool `json:"success,omitempty"`
	Data    struct {
		Type       string `json:"type"`
		Attributes struct {
			Sender_id int `json:"sender_id" validate:"required"`
		} `json:"attributes"`
	} `json:"data"`
}
