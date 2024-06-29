package helper

import (
	"errors"
	"net/http"
	"strings"

	"github.com/go-chi/jwtauth/v5"
)

func CreateJWTToken(username string) (string, error) {
	jwtSecret := GetEnv("JWT_SECRET")
	tokenAuth := jwtauth.New("HS256", []byte(jwtSecret), nil)

	_, tokenString, err := tokenAuth.Encode(map[string]interface{}{"user_id": username})
	return tokenString, err
}

func DecodeJWTToken(JWTToken string) (string, error) {
	jwtSecret := GetEnv("JWT_SECRET")
	tokenAuth := jwtauth.New("HS256", []byte(jwtSecret), nil)

	decodedToken, err := tokenAuth.Decode(JWTToken)
	println(decodedToken, err)
	return "naijason", nil
}

// GetAPIKey extracts an API key from the headers of an HTTP request
// Example:
// Authorization: ApiKey {insert APIkey here}
func GetBearerToken(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authorization key found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("invalid authorization key")
	}
	if vals[0] != "Bearer" {
		return "", errors.New("invalid authorization key")
	}
	// if len(vals[1]) != 64 {
	// 	return "", errors.New("invalid authorization key length")
	// }
	return vals[1], nil
}
