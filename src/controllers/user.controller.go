package controllers
import (
	"context"
	"net/http"
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"api_compiladores/src/models"
	"api_compiladores/src/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


var (
	// UserCollection es la colección de usuarios en MongoDB
	userCollection *mongo.Collection
	regexEmail = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$` // Expresión regular para validar correos electrónicos
)

func SetUserCollection(collection *mongo.Collection) {
	userCollection = collection
}


func isValidEmail(email string) bool {
	// Compilar la expresión regular
	re := regexp.MustCompile(regexEmail)
	// Validar el correo electrónico
	return re.MatchString(email)
}

// funcion para validar los datos del usuario
func validarUsuario(user *models.User) error {
	// Validar los campos del usuario
	if user.Username == "" {
		return fmt.Errorf("el nombre de usuario es obligatorio")
	}
	if user.Password == "" {
		return fmt.Errorf("la contraseña es obligatoria")
	}
	if user.Email == "" {
		return fmt.Errorf("el correo electrónico es obligatorio")
	}
	if len(user.Password) < 6 {
		return fmt.Errorf("la contraseña debe tener al menos 6 caracteres")
	}
	if !isValidEmail(user.Email) {
		return fmt.Errorf("el correo electrónico no es válido")
	}
	return nil
}

func validarData(user *models.User) error {
	// Validar los campos del usuario
	if user.Email == "" {
		return fmt.Errorf("el correo electrónico es obligatorio")
	}
	if user.Password == "" {
		return fmt.Errorf("la contraseña es obligatoria")
	}
	if !isValidEmail(user.Email) {
		return fmt.Errorf("el correo electrónico no es válido")
	}
	return nil
}

func Register(c *gin.Context){
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		sendErrorResponse(c, http.StatusBadRequest, "Datos JSON inválidos", err.Error(), nil)
		return
	}

	if err := validarUsuario(&user); err != nil {
		sendErrorResponse(c, http.StatusBadRequest, "Error de validación", err.Error(), nil)
		return
	}

	user.ID = primitive.NewObjectID()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	// Verificar en base de datos
	err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&user)
	if err == nil {
		sendErrorResponse(c, http.StatusBadRequest,
			fmt.Sprintf("El correo electrónico ya está registrado"),
			nil,
			nil)
		return
	} else if err != mongo.ErrNoDocuments {
		log.Printf("Error al buscar el cliente: %v", err)
		sendErrorResponse(c, http.StatusInternalServerError, "Error interno del servidor", nil, nil)
		return
	}

	user.Password, err = models.EncriptarPassword(user.Password)
	if err != nil {
		log.Printf("Error al encriptar la contraseña: %v", err)
		sendErrorResponse(c, http.StatusInternalServerError, "Error al encriptar la contraseña", nil, nil)
		return
	}

	if _, err := userCollection.InsertOne(ctx, user); err != nil {
		log.Printf("Error al insertar el usuario: %v", err)
		sendErrorResponse(c, http.StatusInternalServerError, "Error al registrar el usuario", nil, nil)
		return
	}

	// Aquí iría la lógica para registrar un nuevo usuario
	c.JSON(http.StatusOK, gin.H{"message": "Registro exitoso"})
}

func Login(c *gin.Context){
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		sendErrorResponse(c, http.StatusBadRequest, "Datos JSON inválidos", err.Error(), nil)
		return
	}

	if err := validarData(&user); err != nil {
		sendErrorResponse(c, http.StatusBadRequest, "Error de validación", err.Error(), nil)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	// Verificar en base de datos
	var foundUser models.User
	err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			sendErrorResponse(c, http.StatusUnauthorized, "Usuario no encontrado", nil, nil)
			return
		}
		log.Printf("Error al buscar el usuario: %v", err)
		sendErrorResponse(c, http.StatusInternalServerError, "Error interno del servidor", nil, nil)
		return
	}

	if err := models.VerificarPassword(foundUser.Password, user.Password); err != nil {
		sendErrorResponse(c, http.StatusUnauthorized, "Credenciales inválidas", nil, nil)
		return
	}

	var token string
	token, err = jwt.GenerarJWT(foundUser.Username)
	if err != nil {
		log.Printf("Error al generar el token: %v", err)
		sendErrorResponse(c, http.StatusInternalServerError, "Error al generar el token", nil, nil)
		return
	}

	// Aquí iría la lógica para iniciar sesión
	c.JSON(http.StatusOK, gin.H{"message": "Inicio de sesión exitoso", "token": token})
}
