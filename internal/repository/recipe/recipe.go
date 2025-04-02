package recipe

import (
	"god/internal/entity"
	"gorm.io/gorm"
)

type Implement struct {
	db *gorm.DB
}

func NewRecipeRepository(db *gorm.DB) *Implement {
	return &Implement{db: db}
}

func (u *Implement) GetById(id int) (*entity.Recipe, error) {
	var recipe *entity.Recipe
	return recipe, u.db.First(&recipe, "id = ?", id).Error
}

func (u *Implement) List(limit, offset int) ([]*entity.Recipe, int, error) {
	var recipes []*entity.Recipe
	var count int64
	err := u.db.Limit(limit).Offset(offset).Find(&recipes).Error
	if err != nil {
		return nil, 0, err
	}

	if err = u.db.Model(&entity.Recipe{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	return recipes, int(count), err
}

func (u *Implement) Create(recipe *entity.Recipe) error {
	return u.db.Create(recipe).Error
}

func (u *Implement) Update(recipe *entity.Recipe) error {
	return u.db.Save(recipe).Error
}

func (u *Implement) Delete(id int) error {
	return u.db.Delete(&entity.User{Id: id}).Error
}
