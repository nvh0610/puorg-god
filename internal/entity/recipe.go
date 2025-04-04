package entity

import "time"

type Recipe struct {
	Id          int       `json:"id" gorm:"id"`
	Title       string    `json:"title" gorm:"title"`
	Description string    `json:"description" gorm:"description"`
	ImageUrl    string    `json:"image_url" gorm:"image_url"`
	Cuisine     string    `json:"cuisine" gorm:"cuisine"` // Loại thực (Việt, Hàn,Âu, v.v.)
	CreatedBy   int       `json:"create_by" gorm:"created_by"`
	CreatedAt   time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"updated_at"`
}

func (r *Recipe) TableName() string {
	return "recipes"
}
