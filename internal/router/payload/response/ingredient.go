package response

import (
	"god/internal/entity"
	"time"
)

type ListIngredientResponse struct {
	Ingredients []*DetailIngredientResponse `json:"ingredients"`
	PaginationResponse
}

type DetailIngredientResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToDetailIngredientResponse(ingredient *entity.Ingredient) *DetailIngredientResponse {
	return &DetailIngredientResponse{
		Id:        ingredient.Id,
		Name:      ingredient.Name,
		CreatedAt: ingredient.CreatedAt,
		UpdatedAt: ingredient.UpdatedAt,
	}
}

func ToListIngredientResponse(ingredients []*entity.Ingredient) []*DetailIngredientResponse {
	var response []*DetailIngredientResponse
	for _, ingredient := range ingredients {
		response = append(response, ToDetailIngredientResponse(ingredient))
	}
	return response
}
