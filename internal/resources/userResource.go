package resources

type UserResource struct {
	Success bool `json:"success,omitempty"`
	Data    struct {
		Id       int    `json:"id"`
		Username string `json:"username"`
	} `json:"data"`
}
