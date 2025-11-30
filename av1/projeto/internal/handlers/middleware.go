package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddlewareGin é o middleware adaptado para Gin
func AuthMiddlewareGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Request.Cookie("user")
		if err != nil || cookie.Value == "" {
			// Redireciona para login
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort() // interrompe a execução do handler
			return
		}

		// Continua para o próximo handler
		c.Next()
	}
}
