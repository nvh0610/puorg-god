package jwt

import (
	"errors"
	"god/pkg/config"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	errExpired      = errors.New("jwt is expired")
	errInvalidToken = errors.New("token is invalid")
)

type Payload struct {
	UserID int    `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

func GenerateToken(id int, role string) (string, error) {
	payload := &Payload{
		UserID: id,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(config.IntEnv("JWT_EXP")) * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(config.StringEnv("JWT_SECRET_KEY")))
}

func Valid(tokenString string) (*Payload, error) {
	payload, err := Parse(KeyFunc(config.StringEnv("JWT_SECRET_KEY")), tokenString)
	if err != nil {
		return nil, err
	}

	return payload, nil
}

func Parse(keyFunc jwt.Keyfunc, tokenString string) (*Payload, error) {
	var payload Payload

	token, err := jwt.ParseWithClaims(tokenString, &payload, keyFunc)
	if err != nil {
		var valErr *jwt.ValidationError

		if errors.As(err, &valErr) {
			if valErr.Errors == jwt.ValidationErrorExpired {
				return &payload, errExpired
			}
		}

		return &payload, err
	}

	if !token.Valid {
		return &payload, errInvalidToken
	}

	return &payload, nil
}
