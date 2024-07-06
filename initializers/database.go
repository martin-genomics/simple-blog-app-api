package initializers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() *gorm.DB {
	var err error

	// connection_string := "postgresql://postgres:UeN5du6A5T0nLjOb@drearily-fabulous-osprey.data-1.use1.tembo.io:5432/postgres"

	dsn := os.Getenv("DB_URL")
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Something went wrong.", err)
	}

	fmt.Println("DATABASE CONNECTION: ", DB)
	return DB
}
