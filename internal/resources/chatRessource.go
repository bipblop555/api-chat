package resources

type ChatRessource struct {
	Success bool `json:"success,omitempty"`
	Data    struct {
		Sender   int    `json:"sender"`
		Receiver int    `json:"receiver"`
		Message  string `json:"message"`
	} `json:"data"`
}
