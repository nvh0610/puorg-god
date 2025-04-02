package user

import (
	"god/internal/entity"
	"god/internal/router/payload/request"
)

func ToModelCreateEntity(user *request.CreateUserRequest) *entity.User {
	return &entity.User{
		Username: user.Username,
		Password: user.Password,
		Role:     user.Role,
		Email:    user.Email,
	}
}

func ToModelUpdateEntity(req *request.UpdateUserRequest, user *entity.User) *entity.User {
	user.Username = req.Username
	return user
}
