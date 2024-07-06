package main

import (
	"log"
	"myapp/initializers"
	"myapp/models"
)

func init() {
	// fmt.Println("----------")
	initializers.LoadEnvVariables()
	// initializers.ConnectToDB()
}

func main() {

	// var user = User{Name: "Martin Tembo", Username: "martintembo1", Email: "543027", Password: "543027"}
	// init()
	initializers.LoadEnvVariables()
	db := initializers.ConnectToDB()
	err := db.AutoMigrate(&models.User{}, &models.Post{})
	// fmt.Println(err)
	if err != nil {
		log.Fatal(err)
	}

}
