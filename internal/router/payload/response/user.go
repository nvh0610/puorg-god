package response

import (
	"god/internal/entity"
	"time"
)

type DetailUserResponse struct {
	Id        int       `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToDetailUserResponse(user *entity.User) *DetailUserResponse {
	return &DetailUserResponse{
		Id:        user.Id,
		Email:     user.Email,
		Username:  user.Username,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToListUserResponse(users []*entity.User) []*DetailUserResponse {
	var res []*DetailUserResponse
	for _, user := range users {
		res = append(res, ToDetailUserResponse(user))
	}
	return res
}

type ListUserResponse struct {
	PaginationResponse
	Users []*DetailUserResponse `json:"users"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}
