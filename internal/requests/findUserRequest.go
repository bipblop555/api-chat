package requests

type FindUserRequest struct {
	Data struct {
		Type       string `json:"type"`
		Attributes struct {
			Username string `json:"username" validate:"max=55|min=3|required"`
		} `json:"attributes"`
	} `json:"data"`
}
