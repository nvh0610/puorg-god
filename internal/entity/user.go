package entity

import "time"

const (
	USER_ROLE_ADMIN = "admin"
	USER_ROLE_USER  = "user"
)

type User struct {
	Id        int       `json:"id" gorm:"id"`
	Email     string    `json:"email" gorm:"email"`
	Username  string    `json:"username" gorm:"username"`
	Password  string    `json:"password" gorm:"password"`
	Role      string    `json:"role" gorm:"role"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
}

func (u *User) TableName() string {
	return "users"
}
