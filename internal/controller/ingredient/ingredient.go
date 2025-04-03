package ingredient

import (
	customStatus "god/internal/common/error"
	"god/internal/repository"
	"god/internal/router/payload/response"
	"god/pkg/resp"
	"god/pkg/utils"
	"net/http"
)

type IngredientController struct {
	repo repository.Registry
}

func NewIngredientController(ingredientRepo repository.Registry) Controller {
	return &IngredientController{
		repo: ingredientRepo,
	}
}

func (u *IngredientController) ListIngredient(w http.ResponseWriter, r *http.Request) {
	page, limit := utils.SetDefaultPagination(r.URL.Query())
	search := r.URL.Query().Get("name")
	offset := (page - 1) * limit

	ingredients, total, err := u.repo.Ingredient().List(limit, offset, search)
	if err != nil {
		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	data := response.ListIngredientResponse{
		Ingredients: response.ToListIngredientResponse(ingredients),
		PaginationResponse: response.PaginationResponse{
			TotalPage: utils.CalculatorTotalPage(total, limit),
			Limit:     limit,
			Page:      page,
		},
	}

	resp.Return(w, http.StatusOK, customStatus.SUCCESS, data)
}
