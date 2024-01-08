package requests

type StoreUserRequest struct {
	Success bool `json:"success,omitempty"`
	Data    struct {
		Type       string `json:"type"`
		Attributes struct {
			Email    string `json:"email_address" validate:"email|min=6|required"`
			Username string `json:"username" validate:"max=55|min=3|required"`
			Password string `json:"password" validate:"max=255|min=5|required"`
		} `json:"attributes"`
	} `json:"data"`
}
