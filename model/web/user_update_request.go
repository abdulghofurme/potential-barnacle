package web

type UserUpdateRequest struct {
	Id          string `validate:"required"`
	Username    string `validate:"required"`
	PhoneNumber string `validate:"required"`
	Role        string `validate:"required"`
}
