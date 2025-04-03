package recipe

import (
	"errors"
	"github.com/go-chi/chi/v5"
	customStatus "god/internal/common/error"
	"god/internal/entity"
	"god/internal/repository"
	"god/internal/router/payload/request"
	"god/internal/router/payload/response"
	"god/pkg/helper"
	"god/pkg/resp"
	"god/pkg/utils"
	"gorm.io/gorm"
	"net/http"
	"sort"
	"strconv"
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

	if err := u.ValidateInstructions(req.Instructions); err != nil {
		resp.Return(w, http.StatusBadRequest, customStatus.INVALID_PARAMS, err.Error())
		return
	}

	err := u.repo.DoInTx(func(txRepo repository.Registry) error {
		recipeEntity := ToModelCreateEntity(req)
		err := txRepo.Recipe().Create(recipeEntity)
		if err != nil {
			return err
		}

		err = u.createRecipeIngredient(req.Ingredients, recipeEntity.Id, txRepo)
		if err != nil {
			return err
		}

		err = u.createInstruction(req.Instructions, recipeEntity.Id, txRepo)
		if err != nil {
			return err
		}

		return err
	})

	if err != nil {
		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	resp.Return(w, http.StatusOK, customStatus.SUCCESS, nil)
}

func (u *RecipeController) UpdateRecipe(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idInt, _ := strconv.Atoi(id)
	req := &request.UpdateRecipeRequest{}
	if err := utils.BindAndValidate(r, req); err != nil {
		resp.Return(w, http.StatusBadRequest, customStatus.INVALID_PARAMS, err.Error())
		return
	}

	if err := u.ValidateInstructions(req.Instructions); err != nil {
		resp.Return(w, http.StatusBadRequest, customStatus.INVALID_PARAMS, err.Error())
		return
	}

	recipe, err := u.repo.Recipe().GetById(idInt)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.Return(w, http.StatusNotFound, customStatus.RECIPE_NOT_FOUND, nil)
			return
		}

		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	err = u.repo.DoInTx(func(txRepo repository.Registry) error {
		recipeEntity := ToModelUpdateEntity(req, recipe)
		err = txRepo.Recipe().Update(recipeEntity)
		if err != nil {
			return err
		}

		err = txRepo.RecipeIngredient().DeleteByRecipeId(idInt)
		if err != nil {
			return err
		}

		err = u.createRecipeIngredient(req.Ingredients, idInt, txRepo)
		if err != nil {
			return err
		}

		err = txRepo.Instruction().DeleteByRecipeId(idInt)
		if err != nil {
			return err
		}

		err = u.createInstruction(req.Instructions, idInt, txRepo)
		if err != nil {
			return err
		}

		return err
	})

	if err != nil {
		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	resp.Return(w, http.StatusOK, customStatus.SUCCESS, nil)
}

func (u *RecipeController) createRecipeIngredient(ingredientRequest []*request.IngredientRequest, recipeId int, txRepo repository.Registry) error {
	var ingredients []*entity.RecipeIngredient
	var err error
	for _, i := range ingredientRequest {
		ingredient, err := txRepo.Ingredient().GetOrCreate(ToModelIngredientEntity(i))
		if err != nil {
			return err
		}
		var recipeIngredient *entity.RecipeIngredient
		recipeIngredient = &entity.RecipeIngredient{
			RecipeId:     recipeId,
			IngredientId: ingredient.Id,
			Quantity:     i.Quantity,
		}

		ingredients = append(ingredients, recipeIngredient)
	}

	if len(ingredients) > 0 {
		err = txRepo.RecipeIngredient().CreateBatch(ingredients)
		if err != nil {
			return errors.New("failed to create recipe ingredients")
		}
	}

	return err
}

func (u *RecipeController) createInstruction(instructionRequest []*request.InstructionRequest, recipeId int, txRepo repository.Registry) error {
	var err error
	for _, i := range instructionRequest {
		instruction := &entity.Instruction{
			RecipeId: recipeId,
			Step:     i.Step,
			Content:  i.Content,
		}

		err = txRepo.Instruction().Create(instruction)
		if err != nil {
			return errors.New("failed to create instruction")
		}
	}

	return err
}

func (u *RecipeController) GetDistinctCuisines(w http.ResponseWriter, r *http.Request) {
	page, limit := utils.SetDefaultPagination(r.URL.Query())
	search := r.URL.Query().Get("cuisine")
	offset := (page - 1) * limit
	cuisines, total, err := u.repo.Recipe().GetDistinctCuisines(limit, offset, search)
	if err != nil {
		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	data := response.GetDistinctCuisinesResponse{
		Cuisines: cuisines,
		PaginationResponse: response.PaginationResponse{
			TotalPage: utils.CalculatorTotalPage(total, limit),
			Limit:     limit,
			Page:      page,
		},
	}

	resp.Return(w, http.StatusOK, customStatus.SUCCESS, data)
}

func (u *RecipeController) GetListRecipe(w http.ResponseWriter, r *http.Request) {
	page, limit := utils.SetDefaultPagination(r.URL.Query())
	offset := (page - 1) * limit
	searchCuisine := r.URL.Query().Get("cuisine")
	searchTitle := r.URL.Query().Get("title")
	searchIngredients := r.URL.Query().Get("ingredients")

	var ingredients []string
	if searchIngredients != "" {
		ingredients = helper.ToArray(searchIngredients)
	}

	recipes, total, err := u.repo.Recipe().List(limit, offset, searchCuisine, searchTitle, ingredients)
	if err != nil {
		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	data := response.ListRecipeResponse{
		Recipes: response.ToRecipeResponse(recipes),
		PaginationResponse: response.PaginationResponse{
			TotalPage: utils.CalculatorTotalPage(total, limit),
			Limit:     limit,
			Page:      page,
		},
	}

	resp.Return(w, http.StatusOK, customStatus.SUCCESS, data)
}

func (u *RecipeController) GetRecipeById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idInt, _ := strconv.Atoi(id)
	recipe, err := u.repo.Recipe().GetDetailById(idInt)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.Return(w, http.StatusNotFound, customStatus.RECIPE_NOT_FOUND, nil)
			return
		}

		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	resp.Return(w, http.StatusOK, customStatus.SUCCESS, response.ToDetailRecipeResponse(recipe))
}

func (u *RecipeController) DeleteRecipeById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idInt, _ := strconv.Atoi(id)
	_, err := u.repo.Recipe().GetById(idInt)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.Return(w, http.StatusNotFound, customStatus.RECIPE_NOT_FOUND, nil)
			return
		}

		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	err = u.repo.Recipe().Delete(idInt)
	if err != nil {
		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	resp.Return(w, http.StatusOK, customStatus.SUCCESS, nil)
}

func (u *RecipeController) ValidateInstructions(req []*request.InstructionRequest) error {
	stepMap := make(map[int]bool)

	for _, instruction := range req {
		if stepMap[instruction.Step] {
			return errors.New("steps in instructions must be unique")
		}
		stepMap[instruction.Step] = true
	}

	var steps []int
	for step := range stepMap {
		steps = append(steps, step)
	}
	sort.Ints(steps)

	for i := 1; i < len(steps); i++ {
		if steps[i] != steps[i-1]+1 {
			return errors.New("steps in instructions must be in ascending order without gaps")
		}
	}

	return nil
}
