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

type RecipeDTO struct {
	ID                  int                    `json:"id" gorm:"column:id"`
	Title               string                 `json:"title" gorm:"column:title"`
	Cuisine             string                 `json:"cuisine" gorm:"column:cuisine"`
	ImageURL            string                 `json:"image_url" gorm:"column:image_url"`
	CreatedBy           int                    `json:"created_by" gorm:"column:created_by"`
	Ingredients         string                 `json:"ingredients" gorm:"column:ingredients"`
	CreateAt            time.Time              `json:"created_at" gorm:"column:created_at"`
	RecipeIngredientDTO []*RecipeIngredientDTO `json:"recipe_ingredients" gorm:"-"`
}

type RecipeIngredientDTO struct {
	ID       int    `json:"id" gorm:"column:id"`
	Name     string `json:"name" gorm:"column:name"`
	Quantity string `json:"quantity" gorm:"column:quantity"`
}

func (u *Implement) List(limit, offset int, searchCuisine, searchTitle string, searchIngredients []string) ([]*RecipeDTO, int, error) {
	var recipesRaw []*RecipeDTO
	var count int64

	applyFilters := func(db *gorm.DB) *gorm.DB {
		if searchCuisine != "" {
			db = db.Where("r.cuisine LIKE ?", "%"+searchCuisine+"%")
		}
		if searchTitle != "" {
			db = db.Where("r.title LIKE ?", "%"+searchTitle+"%")
		}
		if len(searchIngredients) > 0 {
			db = db.Where(`
				EXISTS (
					SELECT 1 FROM recipe_ingredients ri2
					JOIN ingredients i2 ON ri2.ingredient_id = i2.id
					WHERE ri2.recipe_id = r.id AND i2.name IN (?)
				)
			`, searchIngredients)
		}
		return db
	}

	countQuery := applyFilters(u.db.Table("recipes r"))
	if err := countQuery.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	dataQuery := u.db.Table("recipes r").
		Select(`
			r.id AS id,
			r.title,
			r.cuisine,
			r.created_by,
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
		Group("r.id").
		Limit(limit).
		Offset(offset)

	dataQuery = applyFilters(dataQuery)

	if err := dataQuery.Scan(&recipesRaw).Error; err != nil {
		return nil, 0, err
	}

	var recipes []*RecipeDTO
	for _, recipe := range recipesRaw {
		var ingredients []*RecipeIngredientDTO
		if err := json.Unmarshal([]byte(recipe.Ingredients), &ingredients); err != nil {
			return nil, 0, err
		}
		recipes = append(recipes, &RecipeDTO{
			ID:                  recipe.ID,
			Title:               recipe.Title,
			Cuisine:             recipe.Cuisine,
			ImageURL:            recipe.ImageURL,
			CreatedBy:           recipe.CreatedBy,
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
	CreatedBy           int                     `json:"created_by" gorm:"column:created_by"`
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
			r.created_by,
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
		CreatedBy:           recipeRaw.CreatedBy,
		CreateAt:            recipeRaw.CreateAt,
		UpdateAt:            recipeRaw.UpdateAt,
		RecipeIngredientDTO: ingredients,
		InstructionDTO:      instructions,
	}, nil
}
