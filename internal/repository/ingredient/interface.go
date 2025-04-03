package ingredient

import "god/internal/entity"

type Repository interface {
	GetById(id int) (*entity.Ingredient, error)
	Create(ingredient *entity.Ingredient) error
	Update(ingredient *entity.Ingredient) error
	Delete(id int) error
	List(limit, offset int, search string) ([]*entity.Ingredient, int, error)
	GetOrCreate(ingredient *entity.Ingredient) (*entity.Ingredient, error)
}
