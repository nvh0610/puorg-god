package instruction_image

import (
	"god/internal/entity"

	"gorm.io/gorm"
)

type Implement struct {
	db *gorm.DB
}

func NewInstructionImageRepository(db *gorm.DB) *Implement {
	return &Implement{db: db}
}

func (u *Implement) GetById(id int) (*entity.InstructionImage, error) {
	var instructionImage *entity.InstructionImage
	return instructionImage, u.db.First(&instructionImage, "id = ?", id).Error
}

func (u *Implement) List(limit, offset int) ([]*entity.InstructionImage, int, error) {
	var instructionImages []*entity.InstructionImage
	var count int64
	err := u.db.Limit(limit).Offset(offset).Find(&instructionImages).Error
	if err != nil {
		return nil, 0, err
	}

	if err = u.db.Model(&entity.InstructionImage{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	return instructionImages, int(count), err
}

func (u *Implement) Create(recipe *entity.InstructionImage) error {
	return u.db.Create(recipe).Error
}

func (u *Implement) Update(recipe *entity.InstructionImage) error {
	return u.db.Save(recipe).Error
}

func (u *Implement) Delete(id int) error {
	return u.db.Delete(&entity.InstructionImage{Id: id}).Error
}
