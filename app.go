package main

import (
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "log"
    "os"
    "api_compiladores/src/config"
    "api_compiladores/src/routes"
    "api_compiladores/src/utils"
	"github.com/gin-contrib/cors"
)

func main() {
    // Conectar a Redis
    utils.ConnectRedis()

    err := godotenv.Load()
    if err != nil {
        log.Println("Error cargando el archivo .env:", err)
    }

	// Leer puerto del .env
	port := utils.ValidationENV(os.Getenv("PORT"), "8000")

    uri := utils.ValidationENV(os.Getenv("MONGO_URI"), "mongodb://localhost:27017")


    config.ConnectDB(uri)


	// Luego pasarla a la función que carga los usuarios falsos
	utils.AddManyClientes()
    utils.AddUser()

    r := gin.Default()

    // Habilitar CORS
    r.Use(cors.Default())

    routes.ClienteRoutes(r)
    routes.UserRoutes(r)

    r.Run(":" + port)
}