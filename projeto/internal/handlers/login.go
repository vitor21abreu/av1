package handlers

import (
	"net/http"

	"av1/internal/models" // Importa a estrutura de dados User e a lista de usuários

	"github.com/gin-gonic/gin"
)

// LoginCredentials: Estrutura usada para receber e validar os dados do formulário de login.
// As tags `form:"..."` mapeiam os campos do HTML, e `binding:"required"` força que sejam preenchidos.
type LoginCredentials struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

// GetLoginPage: Handler para exibir a página de login.
// Analisa parâmetros de query (?error=...) para exibir mensagens de erro ao usuário.
func GetLoginPage(c *gin.Context) {

	msgErr := c.Query("error")
	var display string

	// Define a mensagem de erro a ser exibida no template
	if msgErr == "invalid" {
		display = "E-mail ou senha incorretos."
	} else if msgErr == "missing" {
		display = "Preencha todos os campos."
	}

	// Renderiza o template de login, passando a mensagem de erro (se houver)
	c.HTML(http.StatusOK, "login.html", gin.H{
		"Error": display,
	})
}

// PostLogin: Handler para processar o envio do formulário de login.
// Realiza a validação dos dados e a autenticação do usuário.
func PostLogin(c *gin.Context) {

	var creds LoginCredentials

	// 1. Valida se os campos 'email' e 'password' foram preenchidos
	if err := c.ShouldBind(&creds); err != nil {
		// Se faltar algum campo, redireciona com o erro 'missing'
		c.Redirect(http.StatusFound, "/login?error=missing")
		return
	}

	var logged *models.User // Variável para armazenar o usuário autenticado

	// 2. Busca o usuário na lista de usuários fake (models.Users)
	for _, user := range models.Users {
		// Verifica se o email e a senha correspondem
		if user.Email == creds.Email && user.Senha == creds.Password {
			logged = &user
			break // Encontrado, interrompe o loop
		}
	}

	// 3. Verifica se o usuário foi encontrado (autenticação falhou)
	if logged == nil {
		// Se credenciais inválidas, redireciona com o erro 'invalid'
		c.Redirect(http.StatusFound, "/login?error=invalid")
		return
	}

	// 4. Autenticação bem-sucedida: Define o cookie de sessão.
	// O email do usuário está sendo usado como token de autenticação.
	// Parâmetros: nome, valor, maxAge (3600s = 1h), path, domain, secure, httpOnly
	c.SetCookie("auth_token", logged.Email, 3600, "/", "", false, true)

	// 5. Redireciona para a página principal
	c.Redirect(http.StatusFound, "/home")
}

// PostLogout: Handler para finalizar a sessão do usuário.
// Remove o cookie de autenticação e redireciona para a página de login.
func PostLogout(c *gin.Context) {

	// Remove o cookie de sessão. Definir MaxAge para -1 força o navegador a deletá-lo.
	c.SetCookie("auth_token", "", -1, "/", "", false, true)

	// Redireciona para a tela de login
	c.Redirect(http.StatusFound, "/login")
}
