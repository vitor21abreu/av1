package handlers

import (
	"av1/internal/models" // Importa a definição da estrutura User
	"net/http"            // Pacote padrão HTTP

	"github.com/gin-gonic/gin" // Importa o framework Gin
)

// Helper Function: Extrai o *models.User do gin.Context de forma segura.
func getUserFromContext(c *gin.Context) *models.User {
	user, exists := c.Get("user")

	if !exists {
		// Retorna um ponteiro para um User vazio em caso de falha inesperada,
		// evitando panic se o middleware falhar, embora não deva acontecer.
		return &models.User{}
	}

	// Converte a interface vazia para o tipo correto *models.User.
	return user.(*models.User)
}

// GET /
// GetRootPage: Redireciona para /home.
func GetRootPage(c *gin.Context) {
	// Se a requisição chegou aqui, o middleware AuthRequired já verificou o cookie.
	c.Redirect(http.StatusFound, "/home")
}

// GET /home
// GetHomePage: Handler para a página inicial (Home).
func GetHomePage(c *gin.Context) {
	user := getUserFromContext(c)

	// LINHAS CORRIGIDAS
	c.HTML(http.StatusOK, "home.html", gin.H{
		"Title": "Home",
		"Page":  "home",
		"User":  user, // Passa o objeto User para o template
	})
}

// GET /chat
// GetChatPage: Handler para a página de Chat.
func GetChatPage(c *gin.Context) {
	// A Role é buscada separadamente do contexto.
	role := c.MustGet("role").(string)
	user := getUserFromContext(c)

	// LINHAS CORRIGIDAS
	c.HTML(http.StatusOK, "chat.html", gin.H{
		"Title": "Chat",
		"Page":  "chat",
		"setor": role,
		"User":  user,
	})
}

// GET /usuario
// GetSobrePage: Handler para a página de perfil ou informações do usuário.
func GetSobrePage(c *gin.Context) {

	// Acessa o objeto User tipado e seguro.
	user := getUserFromContext(c)

	// LINHAS CORRIGIDAS
	c.HTML(http.StatusOK, "usuario.html", gin.H{
		"Title": "Perfil do Usuário",
		"User":  user, // Passa o objeto User
	})
}

// GET /workspace
// GetProcessosPage: Handler para a página de Área de Trabalho/Processos.
func GetProcessosPage(c *gin.Context) {
	user := getUserFromContext(c)

	// LINHAS CORRIGIDAS
	c.HTML(http.StatusOK, "workspace.html", gin.H{
		"Title": "Área de Trabalho",
		"Page":  "processos",
		"User":  user, // Passa o objeto User para a sidebar
	})
}
