package main

import (
	"server/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Registra todas as rotas de checklist
	routes.RegisterChecklistRoutes(e)
	// Registra todas as rotas de note
	routes.RegisterNotaRoutes(e)

	// Inicia o servidor na porta 8080
	e.Logger.Fatal(e.Start(":8080"))
}