package controller

import (
	"github.com/redis/go-redis/v9"
	"god/internal/controller/auth"
	"god/internal/controller/user"
	"god/internal/repository"
)

type RegistryController struct {
	UserCtrl user.Controller
	AuthCtrl auth.Controller
}

func NewRegistryController(repo repository.Registry, redis *redis.Client) *RegistryController {
	return &RegistryController{
		UserCtrl: user.NewUserController(repo),
		AuthCtrl: auth.NewAuthController(repo, redis),
	}
}
