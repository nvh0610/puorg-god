package recipe

import "net/http"

type Controller interface {
	CreateRecipe(w http.ResponseWriter, r *http.Request)
}
