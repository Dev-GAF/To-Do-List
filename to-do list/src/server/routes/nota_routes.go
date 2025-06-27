package routes

import (
	"server/controllers"
	"github.com/labstack/echo/v4"
)

func RegisterNotaRoutes(e *echo.Echo) {
	nota := e.Group("/notas")

	nota.GET("/:id", controllers.ObterNotaPorID)
	nota.GET("", controllers.ObterNotas)
	nota.POST("", controllers.CriarNota)
	nota.PUT("/:id", controllers.AtualizarNota)
	nota.DELETE("/:id", controllers.RemoverNota)
}
