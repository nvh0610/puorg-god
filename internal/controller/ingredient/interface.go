package ingredient

import "net/http"

type Controller interface {
	ListIngredient(w http.ResponseWriter, r *http.Request)
}
