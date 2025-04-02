package user

import "god/internal/entity"

func IsValidUserRole(role string) bool {
	return role == entity.USER_ROLE_USER
}

func IsValidRole(role string) bool {
	return role == entity.USER_ROLE_USER || role == entity.USER_ROLE_ADMIN
}

func IsValidAdminRole(role string) bool {
	return role == entity.USER_ROLE_ADMIN
}
