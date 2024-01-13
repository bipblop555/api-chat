package resources

type ChatRessource struct {
	Success bool `json:"success,omitempty"`
	Data    struct {
		Sender_id   int    `json:"sender_id"`
		Receiver_id int    `json:"receiver_id"`
		Message     string `json:"message"`
	} `json:"data"`
}
