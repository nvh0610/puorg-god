package ingredient

import (
	"god/internal/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Implement struct {
	db *gorm.DB
}

func NewIngredientRepository(db *gorm.DB) *Implement {
	return &Implement{db: db}
}

func (u *Implement) GetById(id int) (*entity.Ingredient, error) {
	var ingredient *entity.Ingredient
	return ingredient, u.db.First(&ingredient, "id = ?", id).Error
}

func (u *Implement) List(limit, offset int, search string) ([]*entity.Ingredient, int, error) {
	var ingredients []*entity.Ingredient
	var count int64

	baseQuery := u.db.Model(&entity.Ingredient{})
	if search != "" {
		baseQuery = baseQuery.Where("name LIKE ?", "%"+search+"%")
	}

	if err := baseQuery.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := baseQuery.Limit(limit).Offset(offset).Find(&ingredients).Error; err != nil {
		return nil, 0, err
	}

	return ingredients, int(count), nil
}

func (u *Implement) GetOrCreate(ingredient *entity.Ingredient) (*entity.Ingredient, error) {
	err := u.db.Clauses(clause.OnConflict{DoNothing: true}).Create(ingredient).Error
	if err != nil {
		return nil, err
	}

	existing := &entity.Ingredient{}
	if err := u.db.Where("name = ?", ingredient.Name).First(existing).Error; err != nil {
		return nil, err
	}

	return existing, nil
}

func (u *Implement) Create(recipe *entity.Ingredient) error {
	return u.db.Create(recipe).Error
}

func (u *Implement) Update(recipe *entity.Ingredient) error {
	return u.db.Save(recipe).Error
}

func (u *Implement) Delete(id int) error {
	return u.db.Delete(&entity.Ingredient{Id: id}).Error
}
