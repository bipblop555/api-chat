package resources

type DeleteMessageRessource struct {
	Success bool `json:"success,omitempty"`
	Data    struct {
		Id      int    `json:"id"`
		Message string `json:"message"`
	} `json:"data"`
}
