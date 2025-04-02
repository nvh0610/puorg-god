package utils

import "net/http"

func GetUserIdAndRoleFromContext(r *http.Request) (int, string) {
	return r.Context().Value("user_id").(int), r.Context().Value("role").(string)
}
