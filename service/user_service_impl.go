package service

import (
	"context"
	"database/sql"
	"time"

	"abdulghofur.me/pshamo-go/helper"
	"abdulghofur.me/pshamo-go/model/domain"
	"abdulghofur.me/pshamo-go/model/web"
	"abdulghofur.me/pshamo-go/repository"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func NewUserService(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
	}
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func (service *UserServiceImpl) Create(ctx context.Context, userRequest web.UserCreateRequest) web.UserResponse {
	err := service.Validate.Struct(userRequest)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	existingUsers := service.UserRepository.FindByUsernameAndPhoneNumber(ctx, tx, userRequest.Username, userRequest.PhoneNumber)
	if len(existingUsers) > 0 {
		panic("username atau phone number sudah digunakan")
	}
	passwordHash, err := helper.HashPassword(userRequest.Password)
	helper.PanicIfError(err)

	user := domain.User{
		Id:           uuid.NewString(),
		Username:     userRequest.Username,
		PhoneNumber:  userRequest.PhoneNumber,
		Role:         userRequest.Role,
		PasswordHash: passwordHash,
	}

	user = service.UserRepository.Create(ctx, tx, user)

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) Update(ctx context.Context, userRequest web.UserUpdateRequest) web.UserResponse {
	err := service.Validate.Struct(userRequest)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userRequest.Id)
	helper.PanicIfError(err)
	if user.DeletedAt.Valid {
		panic("user tidak lagi aktif")
	}

	existingUsers := service.UserRepository.FindByUsernameAndPhoneNumber(ctx, tx, userRequest.Username, userRequest.PhoneNumber)
	if (len(existingUsers) == 1 && existingUsers[0].Id != user.Id) || len(existingUsers) > 1 {
		panic("username atau phone number sudah digunakan")
	}

	user.UpdatedAt = time.Now()
	user = service.UserRepository.Update(ctx, tx, user)
	return helper.ToUserResponse(user)

}

func (service *UserServiceImpl) Delete(ctx context.Context, userId string) web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userId)
	helper.PanicIfError(err)
	if user.DeletedAt.Valid {
		panic("user tidak lagi aktif")
	}

	user.DeletedAt = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}

	service.UserRepository.Delete(ctx, tx, user)

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) FindById(ctx context.Context, userId string) web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userId)
	helper.PanicIfError(err)
	if user.DeletedAt.Valid {
		panic("user tidak lagi aktif")
	}

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) FindAll(ctx context.Context) []web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	users := service.UserRepository.FindAll(ctx, tx)

	var usersResponse []web.UserResponse
	for _, user := range users {
		usersResponse = append(usersResponse, helper.ToUserResponse(user))
	}

	return usersResponse
}
