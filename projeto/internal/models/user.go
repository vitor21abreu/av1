package models

// User: Define a estrutura de dados de um usuário no sistema.
// Os campos devem ter a primeira letra maiúscula (Exported) para serem acessíveis
// pelos templates HTML e handlers do Gin.
type User struct {
	ID    int    // Identificador único do usuário.
	Nome  string // Nome completo do usuário.
	Email string // Email do usuário (usado como token temporário de sessão).
	Senha string // Senha de acesso
	Role  string // Cargo ou Setor do usuário
}

// Esta base é usada para simular o login e a busca de usuário por cookie.
var Users = []User{
	{ID: 1, Nome: "João Dev", Email: "adev@empresa.com", Senha: "1234", Role: "dev"},
	{ID: 2, Nome: "Maria RH", Email: "rh@empresa.com", Senha: "123", Role: "rh"},
	{ID: 3, Nome: "Carlos Financeiro", Email: "fin@empresa.com", Senha: "123", Role: "financeiro"},
}
