// utils/agregarUsers.go
package utils

import (
	"context"
	"fmt"
	"log"
	"time"
	"os"

	"github.com/jaswdr/faker"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"api_compiladores/src/models"
	"api_compiladores/src/config"
)

func AddManyClientes() {
    dbName := ValidationENV(os.Getenv("DB_NAME"), "movilesdb")
	// Primero obtener la colección
	collection := config.GetCollection(dbName, "clientes")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	count, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		log.Fatal("Error al contar clientes:", err)
	}

	if count > 0 {
		fmt.Println("Ya existen clientes en la base de datos.")
		return
	}

	fmt.Println("Insertando clientes...")
	f := faker.New()
	var clientes []interface{}

	for i := 1; i <= 1000000; i++ {
		cliente := models.Cliente{
			ID:            primitive.NewObjectID(),
			Clave_Cliente: fmt.Sprintf("%010d", i),
			Nombre:        f.Person().Name(),
			Celular:       f.Phone().Number(),
			Email:         f.Internet().Email(),
			Character_Icon: f.IntBetween(0, 9), // Asumiendo que los iconos van del 0 al 10
		}

		clientes = append(clientes, cliente)

		// Inserta en lotes de 1000
		if i%1000 == 0 {
			_, err := collection.InsertMany(ctx, clientes)
			if err != nil {
				log.Fatal("Error al insertar clientes:", err)
			}
			clientes = clientes[:0] // Reiniciar slice
			fmt.Printf("Insertados %d clientes...\n", i)
		}
	}
}
