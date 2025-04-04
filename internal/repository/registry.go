package repository

import (
	"god/internal/repository/ingredient"
	"god/internal/repository/instruction"
	"god/internal/repository/recipe"
	"god/internal/repository/recipe_ingredient"
	"god/internal/repository/user"

	"gorm.io/gorm"
)

type Registry interface {
	User() user.Repository
	Recipe() recipe.Repository
	Ingredient() ingredient.Repository
	RecipeIngredient() recipe_ingredient.Repository
	Instruction() instruction.Repository
	DoInTx(txFunc func(txRepo Registry) error) error
}

type mysqlImplement struct {
	db                   *gorm.DB
	userRepo             user.Repository
	recipeRepo           recipe.Repository
	ingredientRepo       ingredient.Repository
	recipeIngredientRepo recipe_ingredient.Repository
	instructionRepo      instruction.Repository
}

func (m *mysqlImplement) User() user.Repository {
	return m.userRepo
}

func (m *mysqlImplement) Recipe() recipe.Repository {
	return m.recipeRepo
}

func (m *mysqlImplement) Ingredient() ingredient.Repository {
	return m.ingredientRepo
}

func (m *mysqlImplement) RecipeIngredient() recipe_ingredient.Repository {
	return m.recipeIngredientRepo
}

func (m *mysqlImplement) Instruction() instruction.Repository {
	return m.instructionRepo
}

func NewRegistryRepo(db *gorm.DB) Registry {
	return &mysqlImplement{
		db:                   db,
		userRepo:             user.NewUserRepository(db),
		recipeRepo:           recipe.NewRecipeRepository(db),
		ingredientRepo:       ingredient.NewIngredientRepository(db),
		recipeIngredientRepo: recipe_ingredient.NewRecipeIngredientRepository(db),
		instructionRepo:      instruction.NewInstructionRepository(db),
	}
}

func (m *mysqlImplement) DoInTx(txFunc func(txRepo Registry) error) error {
	tx := m.db.Begin()
	txRepo := &mysqlImplement{
		db:                   m.db,
		userRepo:             user.NewUserRepository(tx),
		recipeRepo:           recipe.NewRecipeRepository(tx),
		ingredientRepo:       ingredient.NewIngredientRepository(tx),
		recipeIngredientRepo: recipe_ingredient.NewRecipeIngredientRepository(tx),
		instructionRepo:      instruction.NewInstructionRepository(tx),
	}

	err := txFunc(txRepo)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
