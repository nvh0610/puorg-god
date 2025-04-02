package user

import "god/internal/entity"

type Repository interface {
	GetById(id int) (*entity.User, error)
	Create(user *entity.User) error
	Update(user *entity.User) error
	Delete(id int) error
	CheckExistsByEmail(email string) (bool, error)
	List(limit, offset int) ([]*entity.User, int, error)
	GetByEmail(email string) (*entity.User, error)
}
