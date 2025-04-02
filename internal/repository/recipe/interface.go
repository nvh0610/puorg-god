package recipe

import "god/internal/entity"

type Repository interface {
	GetById(id int) (*entity.Recipe, error)
	Create(recipe *entity.Recipe) error
	Update(recipe *entity.Recipe) error
	Delete(id int) error
	List(limit, offset int) ([]*entity.Recipe, int, error)
}
