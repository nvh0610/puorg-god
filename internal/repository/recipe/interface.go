package recipe

import "god/internal/entity"

type Repository interface {
	GetById(id int) (*entity.Recipe, error)
	Create(recipe *entity.Recipe) error
	Update(recipe *entity.Recipe) error
	Delete(id int) error
	List(limit, offset int, searchCuisine, searchTitle string, searchIngredients []string) ([]*RecipeDTO, int, error)
	GetDistinctCuisines(limit, offset int, search string) ([]string, int, error)
	GetDetailById(id int) (*DetailRecipeDTO, error)
}
