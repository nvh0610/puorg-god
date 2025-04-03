package response

import (
	"god/internal/repository/recipe"
	"time"
)

type GetDistinctCuisinesResponse struct {
	Cuisines []string `json:"cuisines"`
	PaginationResponse
}

type ListRecipeResponse struct {
	Recipes []*RecipeResponse `json:"recipes"`
	PaginationResponse
}

type RecipeResponse struct {
	Id          int                         `json:"id"`
	Title       string                      `json:"title"`
	ImageUrl    string                      `json:"image_url"`
	Cuisine     string                      `json:"cuisine"`
	CreatedAt   time.Time                   `json:"created_at"`
	Ingredients []*DetailIngredientResponse `json:"ingredients"`
}

func ToRecipeResponse(recipe []*recipe.RecipeWithIngredients) []*RecipeResponse {
	var recipes []*RecipeResponse
	for _, r := range recipe {
		recipes = append(recipes, &RecipeResponse{
			Id:          r.ID,
			Title:       r.Title,
			ImageUrl:    r.ImageURL,
			Cuisine:     r.Cuisine,
			CreatedAt:   r.CreateAt,
			Ingredients: ToRecipeIngredientResponse(r.RecipeIngredientDTO),
		})
	}
	return recipes
}

func ToRecipeIngredientResponse(ingredients []*recipe.RecipeIngredientDTO) []*DetailIngredientResponse {
	var res []*DetailIngredientResponse
	for _, i := range ingredients {
		res = append(res, &DetailIngredientResponse{
			Id:   i.ID,
			Name: i.Name,
		})
	}
	return res
}

type DetailRecipeResponse struct {
	Id           int                          `json:"id"`
	Title        string                       `json:"title"`
	Description  string                       `json:"description"`
	ImageUrl     string                       `json:"image_url"`
	Cuisine      string                       `json:"cuisine"`
	CreatedAt    time.Time                    `json:"created_at"`
	UpdatedAt    time.Time                    `json:"updated_at"`
	Ingredients  []*DetailIngredientResponse  `json:"ingredients"`
	Instructions []*DetailInstructionResponse `json:"instructions"`
}

func ToDetailRecipeResponse(recipe *recipe.DetailRecipeDTO) *DetailRecipeResponse {
	return &DetailRecipeResponse{
		Id:           recipe.ID,
		Title:        recipe.Title,
		Description:  recipe.Description,
		ImageUrl:     recipe.ImageURL,
		Cuisine:      recipe.Cuisine,
		CreatedAt:    recipe.CreateAt,
		UpdatedAt:    recipe.UpdateAt,
		Ingredients:  ToRecipeIngredientResponse(recipe.RecipeIngredientDTO),
		Instructions: ToDetailInstructionResponse(recipe.InstructionDTO),
	}
}

func ToDetailInstructionResponse(instructions []*recipe.DetailInstructionDTO) []*DetailInstructionResponse {
	var res []*DetailInstructionResponse
	for _, i := range instructions {
		res = append(res, &DetailInstructionResponse{
			Id:      i.ID,
			Step:    i.Step,
			Content: i.Content,
		})
	}
	return res
}
