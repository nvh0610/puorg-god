package request

type CreateUserRequest struct {
	Email    string `json:"email" validate:"required"`
	Username string `json:"username" validate:"required"`
	Role     string `json:"role" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UpdateUserRequest struct {
	Username string `json:"username" validate:"required"`
}

type UpdateRoleRequest struct {
	Role string `json:"role" validate:"required"`
}

type LoginUserRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required"`
}

type ForgetPasswordRequest struct {
	Email string `json:"email" validate:"required"`
}

type VerifyOtpRequest struct {
	Email string `json:"email" validate:"required"`
	Otp   string `json:"otp" validate:"required"`
}

type ResetPasswordRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
