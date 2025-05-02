package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	username string `json:"username"`
	ID       int    `json:"id"`
	IsAdmin  bool   `json:"is_admin"`
	jwt.StandardClaims
}

func GenerateToken(username string, ID int, isAdmin bool)(string, error){
	secretKey := os.Getenv("SECRET_KEY")
	expritationTime := time.Now().Add(time.Hour * 24).Unix()

	claims := &Claims{
		username: username,
		ID:       ID,
		IsAdmin:  isAdmin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt:  expritationTime,
			IssuedAt: time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return signedToken, nil
}



func ValidationToken(tokenString *string) (*Claims, error) {
	secretKey := os.Getenv("SECRET_KEY")
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(*tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}
 
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}