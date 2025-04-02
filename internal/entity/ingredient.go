package entity

import "time"

type Ingredient struct {
	Id        int       `json:"id" gorm:"id"`
	Name      string    `json:"name" gorm:"name"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
}

func (i *Ingredient) TableName() string {
	return "ingredients"
}
