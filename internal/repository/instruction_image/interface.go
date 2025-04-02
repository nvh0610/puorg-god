package instruction_image

import "god/internal/entity"

type Repository interface {
	GetById(id int) (*entity.InstructionImage, error)
	Create(instructionImage *entity.InstructionImage) error
	Update(instructionImage *entity.InstructionImage) error
	Delete(id int) error
	List(limit, offset int) ([]*entity.InstructionImage, int, error)
	CreateBatch(recipes []*entity.InstructionImage) error
}
