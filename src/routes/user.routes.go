package routes

import (
	"os"
    "github.com/gin-gonic/gin"
    "api_compiladores/src/controllers"
    "api_compiladores/src/config"
	"api_compiladores/src/utils"
)

func UserRoutes(router *gin.Engine) {
	dbName := utils.ValidationENV(os.Getenv("DB_NAME"), "movilesdb")
	// Primero obtener la colección
	collection := config.GetCollection(dbName, "users")
	controllers.SetUserCollection(collection)

	userGroup := router.Group("/api/users")

	{
		userGroup.POST("/register", controllers.Register)
		userGroup.POST("/login", controllers.Login)
	}
}