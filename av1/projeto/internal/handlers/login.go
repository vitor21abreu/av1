package handlers

import (
	"html/template"
	"net/http"
)

// A struct user e o slice users já vêm do outro arquivo handlers/users.go
// Não precisa declarar novamente

func Loginpage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/login.html")
	if err != nil {
		http.Error(w, "Erro ao carregar template", http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}

func Loginsubmit(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	senha := r.FormValue("senha")

	// Percorre o slice users vindo de outro arquivo
	for _, u := range users {
		if u.EMAIL == email && u.SENHA == senha {

			// Cria cookie sessao
			http.SetCookie(w, &http.Cookie{
				Name:  "user",
				Value: u.EMAIL,
				Path:  "/",
			})

			http.Redirect(w, r, "/home", http.StatusSeeOther)
			return
		}
	}

	// Login inválido
	http.Redirect(w, r, "/login?erro=1", http.StatusSeeOther)
}
