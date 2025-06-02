package utils

import (
	"context"
	"fmt"
	"log"
	"time"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"api_compiladores/src/models"
	"api_compiladores/src/config"
)

func AddUser()  {
    dbName := ValidationENV(os.Getenv("DB_NAME"), "movilesdb")
    password, err := models.EncriptarPassword(ValidationENV(os.Getenv("USER_ADMIN_PASSWORD"), "123456"))
	if err != nil {
		log.Fatalf("Error al encriptar la contraseña: %v", err)
	}
	// Primero obtener la colección
	collection := config.GetCollection(dbName, "users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	count, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		log.Fatal("Error al contar usuarios:", err)
	}

	if count > 0 {
		fmt.Println("Ya existen usuarios en la base de datos.")
		return
	}

	fmt.Println("Insertando usuarios...")
	var user models.User

	user.ID = primitive.NewObjectID()
	user.Username = "Draken"
	user.Password = password
	user.Email = "bs.personal.0001@gmail.com"

	if _, err := collection.InsertOne(ctx, user); err != nil {
		log.Printf("Error al insertar el usuario: %v", err)
		return
	}
	fmt.Printf("Usuario %s insertado correctamente.\n", user.Username)
}