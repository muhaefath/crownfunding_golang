package main

import (
	"golang_project/handler"
	"golang_project/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/golang_project?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("api/v1")
	api.POST("/users", userHandler.RegisterUser)

	router.Run()

	// userInput := user.RegisterUserInput{}
	// userInput.Name = "Test Simpan"
	// userInput.Email = "Efamil Simpan"
	// userInput.Password = "123456"

	// userService.RegisterUser(userInput)

}
