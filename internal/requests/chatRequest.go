package requests

type ChatRequest struct {
	Success bool `json:"success,omitempty"`
	Data    struct {
		Type       string `json:"type"`
		Attributes struct {
			Sender int `json:"sender" validate:"required"`
		} `json:"attributes"`
	} `json:"data"`
}
