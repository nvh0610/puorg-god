package entity

import "time"

type Instruction struct {
	Id        int       `json:"id" gorm:"id"`
	RecipeId  int       `json:"recipe_id" gorm:"recipe_id"`
	Step      int       `json:"step" gorm:"step"`
	Content   string    `json:"content" gorm:"content"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
}

func (i *Instruction) TableName() string {
	return "instructions"
}
