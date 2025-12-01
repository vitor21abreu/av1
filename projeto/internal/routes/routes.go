package routes

import (
	"av1/internal/handlers"
	"av1/internal/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRoutes: Configura e define todas as rotas da aplicação Gin.
// Recebe o engine (r) do Gin como parâmetro.
func SetupRoutes(r *gin.Engine) {

	// LOGIN: Rotas públicas (sem autenticação)
	r.GET("/login", handlers.GetLoginPage) // Exibe o formulário de login (GET)
	r.POST("/login", handlers.PostLogin)   // Processa o envio do formulário (POST)
	r.POST("/logout", handlers.PostLogout) // Finaliza a sessão (limpa o cookie)

	// ÁREA PROTEGIDA: Rotas que exigem um usuário logado.
	// Cria um grupo de rotas que compartilharão os mesmos middlewares.
	authorized := r.Group("/")

	// MIDDLEWARE 1: Verifica se o cookie de autenticação existe.
	// Se falhar, interrompe o processo e redireciona para /login.
	authorized.Use(middleware.AuthRequired())

	// MIDDLEWARE 2: Busca o objeto User e sua Role após a autenticação.
	// Injeta as chaves "user" e "role" no contexto do Gin para uso nos handlers.
	authorized.Use(middleware.UserContextMiddleware())

	{
		// Definição das rotas protegidas (só acessíveis se os middlewares acima passarem)
		authorized.GET("/home", handlers.HomePage)
		authorized.GET("/chat", handlers.GetChatPage)

		// Rotas de Conteúdo
		authorized.GET("/workspace", handlers.GetProcessosPage) // Rota para a Área de Trabalho
		authorized.GET("/usuario", handlers.GetSobrePage)       // Rota para o Perfil/Sobre o Usuário

		// Rota raiz: Redireciona a URL principal (/) para /home se o usuário estiver logado.
		authorized.GET("/", handlers.GetRootPage)
	}
}
