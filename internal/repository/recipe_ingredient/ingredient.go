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

func (u *Implement) CreateBatch(recipeIngredient []*entity.RecipeIngredient) error {
	return u.db.Create(recipeIngredient).Error
}

func (u *Implement) Delete(id int) error {
	return u.db.Delete(&entity.RecipeIngredient{Id: id}).Error
}

func (u *Implement) DeleteByRecipeId(id int) error {
	return u.db.Where("recipe_id = ?", id).Delete(&entity.RecipeIngredient{}).Error
}
