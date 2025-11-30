package internal

import (
	"projeto/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// Servir arquivos estáticos
	r.Static("/static", "./static")

	// Carregar templates
	r.LoadHTMLGlob("templates/*")

	// Rotas públicas
	r.GET("/", handlers.HomePage)
	r.GET("/login", handlers.LoginPage)

	// Rotas protegidas com middleware
	auth := r.Group("/", handlers.AuthMiddlewareGin())
	{
		auth.GET("/home", handlers.HomePage)
		auth.GET("/chat", handlers.ChatPage)
		auth.GET("/processos", handlers.ProcessoPage)
		auth.GET("/etapas", handlers.EtapasPage)
		auth.GET("/sobre", handlers.SobrePage)
	}

	// APIs de usuários
	r.GET("/users", handlers.GetUsers)
	r.POST("/users", handlers.CreateUser)

	return r
}
