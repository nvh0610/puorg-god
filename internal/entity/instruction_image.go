package entity

import "time"

type InstructionImage struct {
	Id            int       `json:"id" gorm:"id"`
	InstructionId int       `json:"instruction_id" gorm:"instruction_id"`
	ImageUrl      string    `json:"image_url" gorm:"image_url"`
	CreatedAt     time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"updated_at"`
}

func (i *InstructionImage) TableName() string {
	return "instruction_images"
}
