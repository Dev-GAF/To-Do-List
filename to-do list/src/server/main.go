package main

import (
	"server/routes"

	"github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware CORS
    e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{"http://localhost:5173"},
        AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
    }))

	// Registra todas as rotas de checklist
	routes.RegisterChecklistRoutes(e)
	// Registra todas as rotas de note
	routes.RegisterNotaRoutes(e)

	// Inicia o servidor na porta 8080
	e.Logger.Fatal(e.Start(":8080"))
}