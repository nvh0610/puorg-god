package recipe_ingredient

import (
	"god/internal/entity"
	"gorm.io/gorm"
)

type Implement struct {
	db *gorm.DB
}

func NewRecipeIngredientRepository(db *gorm.DB) *Implement {
	return &Implement{db: db}
}

func (u *Implement) GetById(id int) (*entity.RecipeIngredient, error) {
	var recipeIngredient *entity.RecipeIngredient
	return recipeIngredient, u.db.First(&recipeIngredient, "id = ?", id).Error
}

func (u *Implement) List(limit, offset int) ([]*entity.RecipeIngredient, int, error) {
	var recipeIngredient []*entity.RecipeIngredient
	var count int64
	err := u.db.Limit(limit).Offset(offset).Find(&recipeIngredient).Error
	if err != nil {
		return nil, 0, err
	}

	if err = u.db.Model(&entity.RecipeIngredient{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	return recipeIngredient, int(count), err
}

func (u *Implement) Create(recipe *entity.RecipeIngredient) error {
	return u.db.Create(recipe).Error
}

func (u *Implement) Update(recipe *entity.RecipeIngredient) error {
	return u.db.Save(recipe).Error
}

func (u *Implement) Delete(id int) error {
	return u.db.Delete(&entity.RecipeIngredient{Id: id}).Error
}
