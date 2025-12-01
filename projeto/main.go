package main

import (
	"av1/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	// O padrão * garante que todos os templates (base e blocos de conteúdo) sejam lidos.
	r.LoadHTMLGlob("templates/*.html")

	// r.LoadHTMLGlob("templates/*")

	// 2. ROTA ESTÁTICA OK: Garante que os arquivos em /static/ são servidos.
	r.Static("/static", "./static")

	routes.SetupRoutes(r)

	r.Run(":8080")
}
