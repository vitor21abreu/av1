package middleware

import (
	"net/http"

	"av1/internal/models" // Assumindo que esta é a importação correta do seu modelo User

	"github.com/gin-gonic/gin"
)

// AuthRequired: Garante que o usuário tem um cookie de sessão válido.
// Se o cookie não existir ou for inválido, redireciona para /login e interrompe a requisição.
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {

		token, err := c.Cookie("auth_token")
		if err != nil || token == "" {
			c.Redirect(http.StatusFound, "/login?error=no_session")
			c.Abort() // Aborta a cadeia de handlers (impede acesso à rota)
			return
		}

		c.Next() // Permite que a requisição siga para o próximo middleware/handler
	}
}

// UserContextMiddleware: Busca o objeto User (usuário logado) e injeta seus dados no contexto do Gin.
// Este middleware deve ser executado APÓS o AuthRequired.
func UserContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := GetLoggedUser(c)

		if user == nil {
			// Se o usuário não for encontrado (ex: sessão expirada ou dados corrompidos)
			c.Redirect(http.StatusFound, "/login?error=invalid_session")
			c.Abort() // Aborta, pois o handler não conseguirá dados do usuário
			return
		}

		// ESSENCIAL: Armazena o objeto User completo e a Role no contexto.
		// Isso permite que os handlers busquem esses dados usando c.MustGet("user") ou c.MustGet("role").
		c.Set("user", user)
		c.Set("role", user.Role)

		c.Next() // Permite que a requisição siga para o handler da página (Ex: GetChatPage)
	}
}

// GetLoggedUser: Busca o objeto User no slice models.Users usando o cookie de autenticação.
// Retorna o ponteiro para o objeto User se encontrado, ou nil.
func GetLoggedUser(c *gin.Context) *models.User {

	token, err := c.Cookie("auth_token")
	if err != nil {
		return nil
	}

	// Itera sobre todos os usuários para encontrar a correspondência de token/email
	for _, u := range models.Users {
		if token == u.Email { // Assumindo que o token é o Email do usuário
			return &u
		}
	}

	return nil // Usuário não encontrado
}
