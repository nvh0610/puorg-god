package recipe_ingredient

import "god/internal/entity"

type Repository interface {
	Delete(id int) error
	DeleteByRecipeId(id int) error
	CreateBatch(recipeIngredient []*entity.RecipeIngredient) error
}
