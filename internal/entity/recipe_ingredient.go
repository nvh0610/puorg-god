package entity

import "time"

type RecipeIngredient struct {
	Id           int       `json:"id" gorm:"id"`
	RecipeId     int       `json:"recipe_id" gorm:"recipe_id"`
	IngredientId int       `json:"ingredient_id" gorm:"ingredient_id"`
	Quantity     string    `json:"quantity" gorm:"quantity"`
	CreatedAt    time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"updated_at"`
}

func (i *RecipeIngredient) TableName() string {
	return "recipe_ingredients"
}
