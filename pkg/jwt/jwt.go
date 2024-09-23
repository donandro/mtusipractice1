package jwt

import (
	"fmt"
	"math/rand"
	"time"

	"pillsreminder/pkg/jwt/models"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(customClaims map[string]interface{}, jwtSecret string, expire time.Time) (string, error) {
	claims := jwt.MapClaims{}

	for k, v := range customClaims {
		claims[k] = v
	}

	claims["exp"] = expire.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ValidateToken(tokenStr string, jwtSecret string) bool {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})

	if err != nil || !token.Valid {
		return false
	}

	return true
}

func GetClaims(tokenStr string, jwtSecret string) (bool, map[string]interface{}, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})

	if err != nil || !token.Valid {
		return false, nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return token.Valid, claims, nil
	}

	return token.Valid, nil, fmt.Errorf("failed to parse claims")
}

func GenerateRefreshToken(expiry time.Time) models.RefreshToken {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	tokenSlice := make([]byte, 50)

	for i := range tokenSlice {
		tokenSlice[i] = charset[rng.Intn(len(charset))]
	}

	return models.RefreshToken{
		Token:   string(tokenSlice),
		Expires: expiry,
	}
}
