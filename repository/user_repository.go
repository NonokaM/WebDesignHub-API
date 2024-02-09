package repository

import "github.com/NonokaM/WebDesignHub-API/model"

type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}
