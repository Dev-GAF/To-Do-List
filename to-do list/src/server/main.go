package main

import (
	"fmt"
	"log"
	"server/database"
	"server/repositories"
	"server/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Inicializa conexão com o MongoDB
	client, err := database.ConexaoBD()
	if err != nil {
		log.Fatalf("Erro ao conectar ao MongoDB: %v", err)
	}
	fmt.Println("Conexão bem-sucedida com MongoDB:")

	db := client.Database("testdb") // nome do banco

	repositories.SetNotaCollection(db.Collection("notas"))
	repositories.SetChecklistCollection(db.Collection("checklists"))

	// Inicia Echo
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
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}