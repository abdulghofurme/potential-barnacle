package web

type UserCreateRequest struct {
	Username    string `validate:"required"`
	PhoneNumber string `validate:"required"`
	Role        string `validate:"required"`
	Password    string `validate:"required"`
}
