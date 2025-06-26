package controllers

import (
	"net/http"			 			// Para códigos de status HTTP (200, 404, etc.)
	// "strconv" 						// Conversão 
	// "time" 							// Para obter timestamps (como ID único)

	"github.com/labstack/echo/v4" 	// Framework Echo
	"server/models" 				// Pacote interno
	"server/repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GET: Listar todos os checklists
func ListarChecklists(c echo.Context) error {
	checklists, err := repositories.ListarChecklists()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Erro ao buscar checklists"})
	}

	if checklists == nil {
		checklists = []models.Checklist{}
	}

	return c.JSON(http.StatusOK, checklists)
}

// POST: Criar novo checklist
func CriarChecklist(c echo.Context) error {
	var checklist models.Checklist
	if err := c.Bind(&checklist); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Dados inválidos"})
	}

	checklist.ID = primitive.NewObjectID()

	for i := range checklist.Items {
		checklist.Items[i].ID = primitive.NewObjectID()
	}

	err := repositories.CriarChecklist(checklist)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Erro ao salvar checklist"})
	}

	return c.JSON(http.StatusCreated, checklist)
}

// PUT: Atualizar checklist pelo ID
func AtualizarChecklist(c echo.Context) error {
	idStr := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "ID inválido"})
	}

	var checklistAtualizado models.Checklist
	if err := c.Bind(&checklistAtualizado); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if checklistAtualizado.ID != objectID {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "ID do corpo diferente do ID da URL"})
	}

	err = repositories.AtualizarChecklist(objectID, checklistAtualizado)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Erro ao atualizar checklist"})
	}

	return c.JSON(http.StatusOK, checklistAtualizado)
}

// DELETE: Remover checklist pelo ID
func RemoverChecklist(c echo.Context) error {
	idStr := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "ID inválido"})
	}

	err = repositories.RemoverChecklist(objectID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Erro ao remover checklist"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Checklist removido com sucesso"})
}
