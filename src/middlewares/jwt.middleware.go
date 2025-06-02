package middlewares

import (
	"api_compiladores/src/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// AuthMiddleware verifica que exista y sea válido un token JWT
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Leer header Authorization: Bearer token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token requerido"})
			c.Abort()
			return
		}

		// Separar "Bearer token"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Formato de token inválido"})
			c.Abort()
			return
		}

		token := parts[1]

		// Verificar token
		claims, err := jwt.ValidarJWT(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido o expirado"})
			c.Abort()
			return
		}

		// Puedes guardar info en el contexto para controladores después
		c.Set("username", claims.Username)

		// Continuar con la petición
		c.Next()
	}
}
