package recipe

import (
	"god/internal/entity"
	"god/internal/router/payload/request"
	"god/pkg/helper"
)

func ToModelCreateEntity(req *request.CreateRecipeRequest) *entity.Recipe {
	return &entity.Recipe{
		Title:       req.Title,
		Description: req.Description,
		ImageUrl:    req.ImageURL,
		Cuisine:     helper.ToLower(req.Cuisine),
	}
}

func ToModelIngredientEntity(req *request.IngredientRequest) *entity.Ingredient {
	return &entity.Ingredient{
		Name: helper.ToLower(req.Name),
	}
}

func ToModelUpdateEntity(req *request.UpdateRecipeRequest, recipe *entity.Recipe) *entity.Recipe {
	recipe.Title = req.Title
	recipe.Description = req.Description
	recipe.ImageUrl = req.ImageURL
	recipe.Cuisine = helper.ToLower(req.Cuisine)
	return recipe
}
