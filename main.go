package main

import (
	"golang_project/auth"
	"golang_project/campaign"
	"golang_project/handler"
	"golang_project/helper"
	"golang_project/transaction"
	"golang_project/user"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
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
	camppaginRepository := campaign.NewRepository(db)
	transactionRepository := transaction.NewRepository(db)

	userService := user.NewService(userRepository)
	campaignService := campaign.NewService(camppaginRepository)
	transactionService := transaction.NewService(transactionRepository, camppaginRepository)

	authService := auth.NewService()
	userHandler := handler.NewUserHandler(userService, authService)
	campaignHandler := handler.NewCampaignHandler(campaignService)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	router := gin.Default()
	router.Use(cors.Default())
	router.Static("/images", "./images")
	api := router.Group("api/v1")

	// users
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/upload_avatar", authMiddleware(authService, userService), userHandler.UploadAvatar)
	api.POST("/user_fetch", authMiddleware(authService, userService), userHandler.FetchUser)

	// campaign
	api.GET("/campagins", campaignHandler.FindCampaigns)
	api.GET("/campaign/:id", campaignHandler.FindCampaign)
	api.POST("/campagins", authMiddleware(authService, userService), campaignHandler.CreateCampaign)
	api.PUT("/campagins/:id", authMiddleware(authService, userService), campaignHandler.UpdateCampaign)
	api.POST("/campagins-images", authMiddleware(authService, userService), campaignHandler.UploadImage)

	// transaction
	api.GET("/campaign/:id/transaction", authMiddleware(authService, userService), transactionHandler.GetCampaignTransactions)
	api.GET("/transaction", authMiddleware(authService, userService), transactionHandler.GetUserTransactions)
	api.POST("/transaction", authMiddleware(authService, userService), transactionHandler.CreateTransaction)

	router.Run()

}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {

			response := helper.APIRespone("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// bearer token
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIRespone("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIRespone("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userId := int(claim["user_id"].(float64))

		user, err := userService.GetUserById(userId)
		if err != nil {
			response := helper.APIRespone("User not found", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("current_user", user)
	}
}
