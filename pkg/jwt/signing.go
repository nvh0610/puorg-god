package jwt

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

// KeyFunc returns the jwt.KeyFunc for the secret.
func KeyFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	}
}
