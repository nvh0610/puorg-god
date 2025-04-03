package response

type DetailInstructionResponse struct {
	Id      int    `json:"id"`
	Step    int    `json:"step"`
	Content string `json:"content"`
}
