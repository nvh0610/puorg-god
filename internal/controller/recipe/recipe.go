package recipe

import (
	"errors"
	customStatus "god/internal/common/error"
	"god/internal/entity"
	"god/internal/repository"
	"god/internal/router/payload/request"
	"god/pkg/resp"
	"god/pkg/utils"
	"net/http"
)

type RecipeController struct {
	repo repository.Registry
}

func NewRecipeController(recipeRepo repository.Registry) Controller {
	return &RecipeController{
		repo: recipeRepo,
	}
}

func (u *RecipeController) CreateRecipe(w http.ResponseWriter, r *http.Request) {
	req := &request.CreateRecipeRequest{}
	if err := utils.BindAndValidate(r, req); err != nil {
		resp.Return(w, http.StatusBadRequest, customStatus.INVALID_PARAMS, err.Error())
		return
	}

	err := u.repo.DoInTx(func(repo repository.Registry) error {
		recipeEntity := ToModelCreateEntity(req)
		err := repo.Recipe().Create(recipeEntity)
		if err != nil {
			return errors.New("failed to create recipe")
		}

		err = u.createRecipeIngredient(req.Ingredients, recipeEntity.Id)
		if err != nil {
			return errors.New("failed to create recipe ingredient")
		}

		err = u.createInstruction(req.Instructions, recipeEntity.Id)
		if err != nil {
			return errors.New("failed to create instruction")
		}

		return err
	})

	if err != nil {
		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	resp.Return(w, http.StatusOK, customStatus.SUCCESS, nil)
}

func (u *RecipeController) createRecipeIngredient(ingredientRequest []request.CreateIngredientRequest, recipeId int) error {
	var ingredients []*entity.RecipeIngredient
	var err error
	for _, i := range ingredientRequest {
		ingredient, _ := u.repo.Ingredient().GetOrCreate(i.Name)
		var recipeIngredient *entity.RecipeIngredient
		recipeIngredient = &entity.RecipeIngredient{
			RecipeId:     recipeId,
			IngredientId: ingredient.Id,
			Quantity:     i.Quantity,
		}

		ingredients = append(ingredients, recipeIngredient)
	}

	if len(ingredients) > 0 {
		err = u.repo.RecipeIngredient().CreateBatch(ingredients)
		if err != nil {
			return errors.New("failed to create recipe ingredients")
		}
	}

	return err
}

func (u *RecipeController) createInstruction(instructionRequest []request.CreateInstructionRequest, recipeId int) error {
	var err error
	for _, i := range instructionRequest {
		instruction := &entity.Instruction{
			RecipeId: recipeId,
			Step:     i.Step,
			Content:  i.Content,
		}

		err = u.repo.Instruction().Create(instruction)
		if err != nil {
			return errors.New("failed to create instruction")
		}

		var instructionImages []*entity.InstructionImage
		for _, image := range i.Images {
			instructionImages = append(instructionImages, &entity.InstructionImage{
				InstructionId: instruction.Id,
				ImageUrl:      image,
			})
		}

		if len(instructionImages) > 0 {
			err = u.repo.InstructionImage().CreateBatch(instructionImages)
			if err != nil {
				return errors.New("failed to create instruction images")
			}
		}
	}

	return err
}
