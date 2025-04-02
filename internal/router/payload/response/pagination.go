package response

type PaginationResponse struct {
	TotalPage int `json:"total_page"`
	Limit     int `json:"limit"`
	Page      int `json:"page"`
}
