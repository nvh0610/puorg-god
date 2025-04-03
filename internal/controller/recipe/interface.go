package recipe

import "net/http"

type Controller interface {
	CreateRecipe(w http.ResponseWriter, r *http.Request)
	GetDistinctCuisines(w http.ResponseWriter, r *http.Request)
	GetListRecipe(w http.ResponseWriter, r *http.Request)
	GetRecipeById(w http.ResponseWriter, r *http.Request)
	DeleteRecipeById(w http.ResponseWriter, r *http.Request)
}
