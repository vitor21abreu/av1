package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Estrutura que representa um usuário no “banco fictício”
type user struct {
	ID    int    `json:"id"`    // ID único do usuário
	NAME  string `json:"name"`  // Nome do usuário
	EMAIL string `json:"email"` // Email do usuário
	SENHA string `json:"senha"` // Senha  do usuario
	CARGO string `json:"cargo"` // Cargo do usuario
}

// “Banco de dados” em memória — slice com alguns usuários pré‑definidos
var users = []user{
	{1, "ana dev", "adev@hotmail.com", "1234", "dev"},
	{2, "joao dev", "jdev@hotmail.com", "1234", "dev"},
	{3, "luis adm", "ladm@hotmail.com", "1234", "adm"},
	{4, "sofia dev", "sadm@hotmail.com", "1234", "adm"},
}

// Função auxiliar para gerar um novo ID — busca o maior ID atual e soma 1
func getNextID() int {
	maxID := 0
	for _, u := range users {
		if u.ID > maxID {
			maxID = u.ID
		}
	}
	return maxID + 1
}

// Handler para retornar a lista de usuários via JSON
func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

// Handler para criar um novo usuário com base em JSON enviado pelo cliente
func CreateUser(c *gin.Context) {
	// Struct temporária para capturar os dados do JSON de entrada
	var data struct {
		Name  string `json:"name"`
		Email string `json:"email"`
		Senha string `json:"senha"`
		Cargo string `json:"cargo"`
	}

	// Tenta fazer bind do JSON recebido para a struct 'data'
	// Se der erro (JSON mal formado, campos faltando, etc.), responde com 400
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	// Cria um novo usuário com ID gerado automaticamente
	newUser := user{
		ID:    getNextID(),
		NAME:  data.Name,
		EMAIL: data.Email,
		SENHA: data.Senha,
		CARGO: data.Cargo,
	}

	// Adiciona esse novo usuário ao slice 'users'
	// A função append retorna uma **nova slice** — por isso fazemos
	// users = append(users, newUser)
	users = append(users, newUser)

	// Retorna uma resposta de sucesso, com o usuário criado
	c.JSON(http.StatusOK, gin.H{
		"message": "Usuário criado com sucesso!",
		"user":    newUser,
	})
}
