package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"main.go/model"
	"time"
)

var jwtKey = []byte("your-secret-key")

// HashPassword hashes the user's password
func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}
func CheckPasswordHash(pas, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pas))
	fmt.Println(err)
	return err == nil
}

func GenerateToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &jwt.StandardClaims{
		Subject:   string(user.ID),
		ExpiresAt: expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
