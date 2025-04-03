package recipe_ingredient

import "god/internal/entity"

type Repository interface {
	GetById(id int) (*entity.RecipeIngredient, error)
	Create(recipeIngredient *entity.RecipeIngredient) error
	Update(recipeIngredient *entity.RecipeIngredient) error
	Delete(id int) error
	DeleteByRecipeId(id int) error
	List(limit, offset int) ([]*entity.RecipeIngredient, int, error)
	CreateBatch(recipeIngredient []*entity.RecipeIngredient) error
}
