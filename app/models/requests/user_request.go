package requests

type UserLoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}