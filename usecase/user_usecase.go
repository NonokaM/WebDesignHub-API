package usecase

import (
	"github.com/NonokaM/WebDesignHub-API/model"
	"github.com/NonokaM/WebDesignHub-API/repository"
)

type IUserUsecase interface {
	SignUp(user *model.User) (mode.UserResponse, error)
	Login(user model.User) (string, error)
}

type userUsecase struct {
	ur repository.IUserRepository
}

func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
	return &userUsecase{ur}
}
