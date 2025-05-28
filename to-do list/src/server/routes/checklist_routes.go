package routes

import (
	"server/controllers"
	"github.com/labstack/echo/v4"
)

func RegisterChecklistRoutes(e *echo.Echo) {
	checklist := e.Group("/checklists")

	checklist.GET("", controllers.ListarChecklists)           // GET /checklists
	checklist.POST("", controllers.CriarChecklist)            // POST /checklists
	checklist.PUT("/:id", controllers.AtualizarChecklist)     // PUT /checklists/:id
	checklist.DELETE("/:id", controllers.RemoverChecklist)    // DELETE /checklists/:id
}