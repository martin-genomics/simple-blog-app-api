package middlware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	fmt.Println("MIDDLWARE WORKED: ", c.Request.URL.Path)
	tokenString, err := c.Cookie("Authorization")
	// fmt.Println("REACHED", tokenString)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusUnauthorized)

		return
	}
	secretKey := os.Getenv("JWT_SECRET")

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret key used for signing
		return []byte(secretKey), nil
	})

	// Handle parsing errors
	if err != nil {
		fmt.Printf("Error parsing token: %v\n", err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Check if the token is valid
	if !token.Valid {
		fmt.Println("Invalid token")
		c.AbortWithStatus(http.StatusUnauthorized)

		return
	}

	// Extract claims from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("Invalid token claims")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	userId := claims["sub"]

	c.Set("userId", userId)
	c.Next()
}
