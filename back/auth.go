package main

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte("uBF8KO3Dmw6yjyyxg7ERrwA28yv3/RWnOVMzWrtcfi8rAoUYU9hytr+lrlux92axOOAWsYG+EFot04V6vCnyEA==")

var validUser = map[string]string{
	"123@teste.com": "123",
}

func CriaJWTToken(user *ReqLogin) (string, error) {
	if pass, exists := validUser[user.Email]; !exists || pass != user.Senha {
		return "", fmt.Errorf("credenciais inválidas")
	}
	claims := &jwt.MapClaims{
		"expiresAt": 15000,
		"userEmail": user.Email,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func ValidaJWTToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de assinatura inesperado: %v", token.Header["alg"])
		}
		return secret, nil
	})
}

func ValidaJWTHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}
		token, err := ValidaJWTToken(tokenString)
		if err != nil || !token.Valid {
			http.Error(w, "permissão negada", http.StatusForbidden)
			return
		}
		claims := token.Claims.(jwt.MapClaims)
		if _, ok := claims["userEmail"]; !ok {
			http.Error(w, "token inválido", http.StatusForbidden)
			return
		}
		next(w, r)
	}
}
