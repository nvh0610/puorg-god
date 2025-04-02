package controller

import (
	"github.com/redis/go-redis/v9"
	"god/internal/controller/auth"
	"god/internal/controller/ingredient"
	"god/internal/controller/recipe"
	"god/internal/controller/user"
	"god/internal/repository"
)

type RegistryController struct {
	UserCtrl       user.Controller
	AuthCtrl       auth.Controller
	IngredientCtrl ingredient.Controller
	RecipeCtrl     recipe.Controller
}

func NewRegistryController(repo repository.Registry, redis *redis.Client) *RegistryController {
	return &RegistryController{
		UserCtrl:       user.NewUserController(repo),
		AuthCtrl:       auth.NewAuthController(repo, redis),
		IngredientCtrl: ingredient.NewIngredientController(repo),
		RecipeCtrl:     recipe.NewRecipeController(repo),
	}
}
