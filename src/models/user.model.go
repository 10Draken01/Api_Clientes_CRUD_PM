package models

import (
	"golang.org/x/crypto/bcrypt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID 		  primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Username  string             `json:"username" bson:"username"`
	Password  string             `json:"password" bson:"password"`
	Email     string             `json:"email" bson:"email"`
}

func EncriptarPassword(password string) (string, error) {
	// El segundo parámetro es el costo (10 está bien para la mayoría)
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(hashedBytes), err
}

func VerificarPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
