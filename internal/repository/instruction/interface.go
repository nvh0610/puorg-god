package instruction

import "god/internal/entity"

type Repository interface {
	CreateBatch(instructions []*entity.Instruction) error
	DeleteByRecipeId(id int) error
	UpdateBatch(recipes []*entity.Instruction) error
}
