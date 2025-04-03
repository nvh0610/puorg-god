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

func (u *Implement) GetById(id int) (*entity.Instruction, error) {
	var instruction *entity.Instruction
	return instruction, u.db.First(&instruction, "id = ?", id).Error
}

func (u *Implement) List(limit, offset int) ([]*entity.Instruction, int, error) {
	var instructions []*entity.Instruction
	var count int64
	err := u.db.Limit(limit).Offset(offset).Find(&instructions).Error
	if err != nil {
		return nil, 0, err
	}

	if err = u.db.Model(&entity.Instruction{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	return instructions, int(count), err
}

func (u *Implement) Create(recipe *entity.Instruction) error {
	return u.db.Create(recipe).Error
}

func (u *Implement) UpdateBatch(recipes []*entity.Instruction) error {
	return u.db.Save(recipes).Error
}

func (u *Implement) DeleteByRecipeId(id int) error {
	return u.db.Where("recipe_id = ?", id).Delete(&entity.Instruction{}).Error
}

func (u *Implement) Update(recipe *entity.Instruction) error {
	return u.db.Save(recipe).Error
}

func (u *Implement) Delete(id int) error {
	return u.db.Delete(&entity.Instruction{Id: id}).Error
}

func (u *Implement) CreateBatch(recipes []*entity.Instruction) error {
	return u.db.Create(recipes).Error
}
