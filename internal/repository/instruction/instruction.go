package instruction

import (
	"god/internal/entity"

	"gorm.io/gorm"
)

type Implement struct {
	db *gorm.DB
}

func NewInstructionRepository(db *gorm.DB) *Implement {
	return &Implement{db: db}
}

func (u *Implement) UpdateBatch(recipes []*entity.Instruction) error {
	return u.db.Save(recipes).Error
}

func (u *Implement) DeleteByRecipeId(id int) error {
	return u.db.Where("recipe_id = ?", id).Delete(&entity.Instruction{}).Error
}

func (u *Implement) CreateBatch(recipes []*entity.Instruction) error {
	return u.db.Create(recipes).Error
}
