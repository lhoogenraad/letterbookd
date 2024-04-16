package utils

import (
	"time"
	"github.com/dgrijalva/jwt-go"
	"fmt"
)

var secretKey = []byte("secretpassword")

// GenerateToken generates a JWT token with the user ID as part of the claims
func GenerateToken(userId uint, username string, email string) (string, error) {
	claims := jwt.MapClaims{}
	claims["userid"] = userId
	claims["username"] = username
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 34210).Unix() // Token valid for 1 hour

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// VerifyToken verifies a token JWT validate 
func VerifyToken(tokenString string) (jwt.MapClaims, error) {
    // Parse the token
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        // Check the signing method
        _, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
            return nil, fmt.Errorf("Invalid signing method")
        }

        return secretKey, nil
    })

	// Check for errors
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Validate the token
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("Invalid token")
}
