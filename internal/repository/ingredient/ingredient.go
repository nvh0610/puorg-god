package ingredient

import (
	"god/internal/entity"
	"gorm.io/gorm"
)

type Implement struct {
	db *gorm.DB
}

func NewIngredientRepository(db *gorm.DB) *Implement {
	return &Implement{db: db}
}

func (u *Implement) GetById(id int) (*entity.Ingredient, error) {
	var ingredient *entity.Ingredient
	return ingredient, u.db.First(&ingredient, "id = ?", id).Error
}

func (u *Implement) List(limit, offset int) ([]*entity.Ingredient, int, error) {
	var ingredients []*entity.Ingredient
	var count int64
	err := u.db.Limit(limit).Offset(offset).Find(&ingredients).Error
	if err != nil {
		return nil, 0, err
	}

	if err = u.db.Model(&entity.Ingredient{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	return ingredients, int(count), err
}

func (u *Implement) Create(recipe *entity.Ingredient) error {
	return u.db.Create(recipe).Error
}

func (u *Implement) Update(recipe *entity.Ingredient) error {
	return u.db.Save(recipe).Error
}

func (u *Implement) Delete(id int) error {
	return u.db.Delete(&entity.Ingredient{Id: id}).Error
}
