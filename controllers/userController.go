package controllers

import (
	"encoding/json"
	"fmt"
	"myapp/initializers"
	"myapp/models"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	var db = initializers.ConnectToDB()

	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Failed to connect",
			"data":    gin.H{},
		})

		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 11)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Failed to hash the password.",
		})
	}

	user := models.User{Email: body.Email, Password: string(hash)}

	result := db.Create(&user)
	// log.Fatal(err, result)

	if result.Error != nil {
		// c.Status(404)
		c.JSON(400, gin.H{
			"success": false,
			"message": "Failed to connect",
			"data":    gin.H{},
		})

	}

	c.JSON(http.StatusOK, gin.H{})

}

func Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Failed to connect",
			"data":    gin.H{},
		})

		return
	}
	var db = initializers.ConnectToDB()

	var user models.User
	result := db.Where(&models.User{Email: body.Email}).First(&user)

	fmt.Println(result)
	if result.Error != nil {
		c.JSON(404, gin.H{
			"success": false,
			"message": "The provided email address does not exist in the system.",
			"data":    gin.H{},
		})

		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)) != nil {

		c.JSON(401, gin.H{
			"success": false,
			"message": "The provided password is incorrect.",
			"data":    gin.H{},
		})

		return

	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": json.Number(strconv.FormatInt(time.Now().Add(time.Hour*time.Duration(1)).Unix(), 10)),
		"iat": json.Number(strconv.FormatInt(time.Now().Unix(), 10)),
	})

	secret_key := os.Getenv("JWT_SECRET")

	tokenString, err := token.SignedString([]byte(secret_key))
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(401, gin.H{
			"success": true,
			"message": "Could not generate the user token.",
			"data": gin.H{
				"user": nil,
			},
		})

		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(200, gin.H{
		"success": true,
		"message": "The login success.",
		"data": gin.H{
			"token": tokenString,
		},
	})

}
