package instruction

import "god/internal/entity"

type Repository interface {
	GetById(id int) (*entity.Instruction, error)
	Create(instruction *entity.Instruction) error
	Update(instruction *entity.Instruction) error
	Delete(id int) error
	List(limit, offset int) ([]*entity.Instruction, int, error)
	CreateBatch(instructions []*entity.Instruction) error
}
