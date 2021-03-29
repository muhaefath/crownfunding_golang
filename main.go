package main

import (
	"fmt"
	"golang_project/auth"
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
	authService := auth.NewService()

	token, err := authService.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyMn0.rguBXk-qieVMLLu5t9U670syjrRvSpVNFr4AuOCrD")
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println("ERROR")
		fmt.Println("ERROR")
	}

	if token.Valid {
		fmt.Println("VALID")
		fmt.Println("VALID")
		fmt.Println("VALID")
	} else {
		fmt.Println("INVALID")
		fmt.Println("INVALID")
		fmt.Println("INVALID")
	}

	fmt.Println(authService.GenerateToken(1))
	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/upload_avatar", userHandler.UploadAvatar)

	router.Run()

}
