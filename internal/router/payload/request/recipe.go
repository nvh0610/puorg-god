package request

type CreateRecipeRequest struct {
	Title        string                `json:"title" validate:"required"`
	Description  string                `json:"description" validate:"required"`
	ImageURL     string                `json:"image_url"`
	Cuisine      string                `json:"cuisine"`
	Ingredients  []*IngredientRequest  `json:"ingredients" validate:"required,dive"`
	Instructions []*InstructionRequest `json:"instructions" validate:"required,dive"`
}

type IngredientRequest struct {
	Name     string `json:"name" validate:"required"`
	Quantity string `json:"quantity" validate:"required"`
}

type InstructionRequest struct {
	Step    int    `json:"step" validate:"required"`
	Content string `json:"content" validate:"required"`
}

type UpdateRecipeRequest struct {
	Title        string                `json:"title" validate:"required"`
	Description  string                `json:"description" validate:"required"`
	ImageURL     string                `json:"image_url"`
	Cuisine      string                `json:"cuisine"`
	Ingredients  []*IngredientRequest  `json:"ingredients" validate:"required,dive"`
	Instructions []*InstructionRequest `json:"instructions" validate:"required,dive"`
}
