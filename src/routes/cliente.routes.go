package routes

import (
    "os"
    "github.com/gin-gonic/gin"
    "api_compiladores/src/controllers"
    "api_compiladores/src/config"
    "api_compiladores/src/utils"
    "api_compiladores/src/middlewares"
)

func ClienteRoutes(router *gin.Engine) {
	dbName := utils.ValidationENV(os.Getenv("DB_NAME"), "movilesdb")
	// Primero obtener la colección
	clienteCollection := config.GetCollection(dbName, "clientes")
    controllers.SetClienteCollection(clienteCollection)

    clienteGroup := router.Group("/api/clientes")
    clienteGroup.Use(middlewares.AuthMiddleware()) // Middleware de autenticación
    {
        clienteGroup.POST("/", controllers.CreateCliente)
        clienteGroup.GET("/page/:page", controllers.GetClientes)
        clienteGroup.GET("/:clave_cliente", controllers.GetCliente)
        clienteGroup.PUT("/:clave_cliente", controllers.UpdateCliente)
        clienteGroup.DELETE("/:clave_cliente", controllers.DeleteCliente)
    }
}