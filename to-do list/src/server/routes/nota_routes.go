package routes

import (
	"server/controllers"
	"github.com/labstack/echo/v4"
)

func RegisterNotaRoutes(e *echo.Echo) {
	nota := e.Group("/notas")

	nota.GET("", controllers.ListarNotas)
	nota.POST("", controllers.CriarNota)
	nota.PUT("/:id", controllers.AtualizarNota)
	nota.DELETE("/:id", controllers.RemoverNota)
}
