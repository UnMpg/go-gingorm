package main

import (
	"fmt"
	"gin-gorm-postgres/initializers"
	"gin-gorm-postgres/models"
	"log"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("?Could not load environment variable")
	}
	fmt.Println(config)
	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
	fmt.Println("? Migration complete")
}
