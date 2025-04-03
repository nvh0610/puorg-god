package recipe

import (
	"encoding/json"
	"god/internal/entity"
	"gorm.io/gorm"
	"time"
)

type Implement struct {
	db *gorm.DB
}

func NewRecipeRepository(db *gorm.DB) *Implement {
	return &Implement{db: db}
}

func (u *Implement) GetById(id int) (*entity.Recipe, error) {
	var recipe *entity.Recipe
	return recipe, u.db.First(&recipe, "id = ?", id).Error
}

type RecipeWithIngredients struct {
	ID                  int                    `json:"id" gorm:"column:id"`
	Title               string                 `json:"title" gorm:"column:title"`
	Cuisine             string                 `json:"cuisine" gorm:"column:cuisine"`
	ImageURL            string                 `json:"image_url" gorm:"column:image_url"`
	Ingredients         string                 `json:"ingredients" gorm:"column:ingredients"`
	CreateAt            time.Time              `json:"created_at" gorm:"column:created_at"`
	RecipeIngredientDTO []*RecipeIngredientDTO `json:"recipe_ingredients" gorm:"-"`
}

type RecipeIngredientDTO struct {
	ID       int    `json:"id" gorm:"column:id"`
	Name     string `json:"name" gorm:"column:name"`
	Quantity string `json:"quantity" gorm:"column:quantity"`
}

func (u *Implement) List(limit, offset int, searchCuisine, searchTitle string, searchIngredients []string) ([]*RecipeWithIngredients, int, error) {
	var recipesRaw []*RecipeWithIngredients
	var count int64

	query := u.db.Table("recipes r").
		Select(`
			r.id AS id,
			r.title,
			r.cuisine,
			r.image_url,
			r.created_at,
			IFNULL(JSON_ARRAYAGG(
				JSON_OBJECT(
					'id', i.id,
					'name', i.name
				)
			), '[]') AS ingredients
		`).
		Joins("LEFT JOIN recipe_ingredients ri ON r.id = ri.recipe_id").
		Joins("LEFT JOIN ingredients i ON ri.ingredient_id = i.id").
		Group("r.id")

	if searchCuisine != "" {
		query = query.Where("r.cuisine LIKE ?", "%"+searchCuisine+"%")
	}
	if searchTitle != "" {
		query = query.Where("r.title LIKE ?", "%"+searchTitle+"%")
	}
	if len(searchIngredients) > 0 {
		query = query.Where("i.name IN (?)", searchIngredients)
	}

	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Limit(limit).Offset(offset).Scan(&recipesRaw).Error; err != nil {
		return nil, 0, err
	}

	var recipes []*RecipeWithIngredients
	for _, recipe := range recipesRaw {
		var ingredients []*RecipeIngredientDTO
		if err := json.Unmarshal([]byte(recipe.Ingredients), &ingredients); err != nil {
			return nil, 0, err
		}
		recipes = append(recipes, &RecipeWithIngredients{
			ID:                  recipe.ID,
			Title:               recipe.Title,
			Cuisine:             recipe.Cuisine,
			ImageURL:            recipe.ImageURL,
			CreateAt:            recipe.CreateAt,
			RecipeIngredientDTO: ingredients,
		})
	}

	return recipes, int(count), nil
}

func (u *Implement) Create(recipe *entity.Recipe) error {
	return u.db.Create(recipe).Error
}

func (u *Implement) Update(recipe *entity.Recipe) error {
	return u.db.Save(recipe).Error
}

func (u *Implement) Delete(id int) error {
	return u.db.Delete(&entity.Recipe{Id: id}).Error
}

func (u *Implement) GetDistinctCuisines(limit, offset int, search string) ([]string, int, error) {
	var cuisines []string
	var countCuisines int64

	baseQuery := u.db.Model(&entity.Recipe{}).Distinct("cuisine")
	if search != "" {
		baseQuery = baseQuery.Where("cuisine LIKE ?", "%"+search+"%")
	}

	if err := baseQuery.Count(&countCuisines).Error; err != nil {
		return nil, 0, err
	}

	if err := baseQuery.Limit(limit).Offset(offset).Find(&cuisines).Error; err != nil {
		return nil, 0, err
	}

	return cuisines, int(countCuisines), nil
}

type DetailRecipeDTO struct {
	ID                  int                     `json:"id" gorm:"column:id"`
	Title               string                  `json:"title" gorm:"column:title"`
	Description         string                  `json:"description" gorm:"column:description"`
	Cuisine             string                  `json:"cuisine" gorm:"column:cuisine"`
	ImageURL            string                  `json:"image_url" gorm:"column:image_url"`
	Ingredients         string                  `json:"ingredients" gorm:"column:ingredients"`
	CreateAt            time.Time               `json:"created_at" gorm:"column:created_at"`
	UpdateAt            time.Time               `json:"updated_at" gorm:"column:updated_at"`
	RecipeIngredientDTO []*RecipeIngredientDTO  `json:"recipe_ingredients" gorm:"-"`
	Instructions        string                  `json:"instructions" gorm:"column:instructions"`
	InstructionDTO      []*DetailInstructionDTO `json:"recipe_instructions" gorm:"-"`
}

type DetailInstructionDTO struct {
	ID      int    `json:"id" gorm:"column:id"`
	Step    int    `json:"step" gorm:"column:step"`
	Content string `json:"content" gorm:"column:content"`
}

func (u *Implement) GetDetailById(id int) (*DetailRecipeDTO, error) {
	var recipeRaw DetailRecipeDTO
	if err := u.db.Table("recipes r").
		Select(`
			r.id AS id,
			r.title,
			r.description,
			r.cuisine,
			r.image_url,
			r.created_at,
			r.updated_at,
			IFNULL(GROUP_CONCAT(DISTINCT CONCAT('{"id":', i.id, ',"name":"', i.name, '","quantity":"', ri.quantity, '"}') SEPARATOR ','), '[]') AS ingredients,
			IFNULL(GROUP_CONCAT(DISTINCT CONCAT('{"id":', ir.id, ',"step":', ir.step, ',"content":"', ir.content, '"}') SEPARATOR ','), '[]') AS instructions
		`).
		Joins("LEFT JOIN recipe_ingredients ri ON r.id = ri.recipe_id").
		Joins("LEFT JOIN ingredients i ON ri.ingredient_id = i.id").
		Joins("LEFT JOIN instructions ir ON r.id = ir.recipe_id").
		Where("r.id = ?", id).
		Group("r.id").
		Find(&recipeRaw).Error; err != nil {
		return nil, err
	}

	var ingredients []*RecipeIngredientDTO
	if err := json.Unmarshal([]byte("["+recipeRaw.Ingredients+"]"), &ingredients); err != nil {
		return nil, err
	}

	var instructions []*DetailInstructionDTO
	if err := json.Unmarshal([]byte("["+recipeRaw.Instructions+"]"), &instructions); err != nil {
		return nil, err
	}

	return &DetailRecipeDTO{
		ID:                  recipeRaw.ID,
		Title:               recipeRaw.Title,
		Description:         recipeRaw.Description,
		Cuisine:             recipeRaw.Cuisine,
		ImageURL:            recipeRaw.ImageURL,
		CreateAt:            recipeRaw.CreateAt,
		UpdateAt:            recipeRaw.UpdateAt,
		RecipeIngredientDTO: ingredients,
		InstructionDTO:      instructions,
	}, nil
}
