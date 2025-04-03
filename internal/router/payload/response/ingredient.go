package response

import (
	"god/internal/entity"
)

type ListIngredientResponse struct {
	Ingredients []*DetailIngredientResponse `json:"ingredients"`
	PaginationResponse
}

type DetailIngredientResponse struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Quantity string `json:"quantity,omitempty"`
}

func ToDetailIngredientResponse(ingredient *entity.Ingredient) *DetailIngredientResponse {
	return &DetailIngredientResponse{
		Id:   ingredient.Id,
		Name: ingredient.Name,
	}
}

func ToListIngredientResponse(ingredients []*entity.Ingredient) []*DetailIngredientResponse {
	var response []*DetailIngredientResponse
	for _, ingredient := range ingredients {
		response = append(response, ToDetailIngredientResponse(ingredient))
	}
	return response
}
