package handlers

import (
	"net/http"

	"av1/internal/middleware" // Importa o pacote de middleware para acessar funções como GetLoggedUser

	"github.com/gin-gonic/gin" // Importa o framework Gin
)

// HomePage: Handler para a rota /home.
// Exibe a página principal da área logada, passando os dados do usuário.
func HomePage(c *gin.Context) {

	// 1. Busca o objeto User logado.
	// **MELHORIA:** Idealmente, deve usar c.MustGet("user") se o UserContextMiddleware estiver ativo,
	// pois evita a leitura duplicada do cookie.
	user := middleware.GetLoggedUser(c)

	// 2. Renderiza o template home.html.
	// Passa o objeto User completo no mapa de dados ("User": user).
	c.HTML(http.StatusOK, "home.html", gin.H{
		"User": user, // O template acessa os dados usando {{ .User.Nome }}, {{ .User.Role }}, etc.
	})
}
