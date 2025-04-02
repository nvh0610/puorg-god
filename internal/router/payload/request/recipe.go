package request

type CreateRecipeRequest struct {
	Title        string                     `json:"title" validate:"required"`
	Description  string                     `json:"description" validate:"required"`
	ImageURL     string                     `json:"image_url"`
	Cuisine      string                     `json:"cuisine"`
	CreatedBy    int                        `json:"created_by" validate:"required"`
	Ingredients  []CreateIngredientRequest  `json:"ingredients" validate:"required,dive"`
	Instructions []CreateInstructionRequest `json:"instructions" validate:"required,dive"`
}

type CreateIngredientRequest struct {
	Name     string `json:"name" validate:"required"`
	Quantity string `json:"quantity" validate:"required"`
}

type CreateInstructionRequest struct {
	Step    int      `json:"step" validate:"required"`
	Content string   `json:"content" validate:"required"`
	Images  []string `json:"images"`
}
