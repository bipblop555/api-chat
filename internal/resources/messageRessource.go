package resources

type MessageRessource struct {
	Success bool `json:"success,omitempty"`
	Data    struct {
		Sender_id int    `json:"sender"`
		Message   string `json:"message"`
	} `json:"data"`
}
