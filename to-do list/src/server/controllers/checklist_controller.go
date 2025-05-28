package controllers

import (
	"net/http"                       // Para códigos de status HTTP (200, 404, etc.)
	"github.com/labstack/echo/v4"    // Framework Echo
	"server/models"                  // Pacote interno
	"time"                           // Para obter timestamps (como ID único)
	"strconv"                        // Conversão de números para string
)

var checklists []models.Checklist

// GET: Listar todos os checklists
func ListarChecklists(c echo.Context) error {
	return c.JSON(http.StatusOK, checklists)
}

// POST: Criar novo checklist
func CriarChecklist(c echo.Context) error {
	var checklist models.Checklist
	if err := c.Bind(&checklist); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	checklist.ID = strconv.FormatInt(time.Now().UnixNano(), 10)
	checklists = append(checklists, checklist)
	return c.JSON(http.StatusCreated, checklist)
}

// PUT: Atualizar checklist pelo ID
func AtualizarChecklist(c echo.Context) error {
	id := c.Param("id")
	var atualizado models.Checklist
	if err := c.Bind(&atualizado); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	for i, cl := range checklists {
		if cl.ID == id {
			atualizado.ID = id
			checklists[i] = atualizado
			return c.JSON(http.StatusOK, atualizado)
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"error": "Checklist não encontrado"})
}

// DELETE: Remover checklist pelo ID
func RemoverChecklist(c echo.Context) error {
	id := c.Param("id")
	for i, cl := range checklists {
		if cl.ID == id {
			checklists = append(checklists[:i], checklists[i+1:]...) // Remove da slice
			return c.JSON(http.StatusOK, map[string]string{"message": "Checklist removido"})
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"error": "Checklist não encontrado"})
}