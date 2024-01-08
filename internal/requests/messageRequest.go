package requests

type SendMessageRequest struct {
	Success bool `json:"success,omitempty"`
	Data    struct {
		Type       string `json:"type"`
		Attributes struct {
			Sender   int    `json:"sender" validate:"required"`
			Receiver int    `json:"receiver" validate:"required"`
			Message  string `json:"message" validate:"required"`
			//CreatedAt *time.Time `json:"created-at" gorm:"type:timestamp"`
			//UpdatedAt *time.Time `json:"updated-at" gorm:"type:timestamp;autoUpdateTime:true"`
		} `json:"attributes"`
	} `json:"data"`
}
