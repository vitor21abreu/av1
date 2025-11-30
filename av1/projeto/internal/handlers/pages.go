package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// rotas
func HomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", nil)
}

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}
