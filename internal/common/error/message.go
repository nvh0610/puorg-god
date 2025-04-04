package error

import (
	"god/pkg/resp"
)

var MsgFlags = map[int]string{
	USER_NOT_FOUND:         "User not found",
	WRONG_PASSWORD:         "Wrong password",
	UPDATE_PASSWORD_FAILED: "Update password failed",
	SUCCESS:                "Success",
	INVALID_PARAMS:         "Invalid params",
	UNAUTHORIZED:           "Unauthorized",
	INTERNAL_SERVER:        "Internal server",
	FORBIDDEN:              "Not allow to process action",
	USER_EXIST:             "User existed",
	CREATE_USER_FAILED:     "Create user failed",
	UPDATE_USER_FAILED:     "Update user failed",
	USER_NOT_ADMIN:         "User not admin",
	WRONG_OTP:              "Wrong otp",
	RECIPE_NOT_FOUND:       "Recipe not found",
}

func InitErrMsg() {
	for key, value := range MsgFlags {
		resp.MsgFlags[key] = value
	}
}
