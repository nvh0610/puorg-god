package user

import (
	"god/internal/entity"
	"gorm.io/gorm"
)

type Implement struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *Implement {
	return &Implement{db: db}
}

func (u *Implement) GetById(id int) (*entity.User, error) {
	var user *entity.User
	return user, u.db.First(&user, "id = ?", id).Error
}

func (u *Implement) GetByEmail(email string) (*entity.User, error) {
	var user *entity.User
	return user, u.db.First(&user, "email = ?", email).Error
}

func (u *Implement) CheckExistsByEmail(email string) (bool, error) {
	var user entity.User
	return user.Id != 0, u.db.First(&user, "email = ?", email).Error
}

func (u *Implement) List(limit, offset int) ([]*entity.User, int, error) {
	var users []*entity.User
	var count int64
	err := u.db.Limit(limit).Offset(offset).Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	if err = u.db.Model(&entity.User{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	return users, int(count), err
}

func (u *Implement) Create(user *entity.User) error {
	return u.db.Create(user).Error
}

func (u *Implement) Update(user *entity.User) error {
	return u.db.Save(user).Error
}

func (u *Implement) Delete(id int) error {
	return u.db.Delete(&entity.User{Id: id}).Error
}
