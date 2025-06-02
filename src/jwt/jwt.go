package jwt

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

func validationENV(env string, envDefault string) string {
    if env == "" {
        return envDefault
    }
    return env
}
var jwtKey = []byte(validationENV(os.Getenv("JWT_KEY"), "miClaveSecreta"))

type Claims struct {
	Username string `json:"Username"`
	jwt.RegisteredClaims
}

func GenerarJWT(username string) (string, error) {
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			Issuer:    "miAppGo",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidarJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}
