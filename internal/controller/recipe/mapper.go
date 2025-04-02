package recipe

import (
	"god/internal/entity"
	"god/internal/router/payload/request"
)

func ToModelCreateEntity(req *request.CreateRecipeRequest) *entity.Recipe {
	return &entity.Recipe{
		Title:       req.Title,
		Description: req.Description,
		ImageUrl:    req.ImageURL,
		Cuisine:     req.Cuisine,
	}
}
