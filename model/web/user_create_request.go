package web

type UserCreateRequest struct {
	Username    string
	PhoneNumber string
	Role        string
	Password    string
}
