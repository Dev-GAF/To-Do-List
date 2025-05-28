package controllers

import (
	"net/http"
	"time"
	"strconv"

	"github.com/labstack/echo/v4"
	"server/models"
)

var notas []models.Nota

// GET: Listar todos as notas
func ListarNotas(c echo.Context) error {
	return c.JSON(http.StatusOK, notas)
}

// POST: Criar nova nota
func CriarNota(c echo.Context) error {
	var nota models.Nota
	if err := c.Bind(&nota); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	nota.ID = int(time.Now().UnixNano() / 1e6) 
	notas = append(notas, nota)
	return c.JSON(http.StatusCreated, nota)
}

// PUT: Atualizar nota pelo ID
func AtualizarNota(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr) 
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	var atualizada models.Nota
	if err := c.Bind(&atualizada); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	for i, n := range notas {
		if n.ID == id {
			atualizada.ID = id
			notas[i] = atualizada
			return c.JSON(http.StatusOK, atualizada)
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"error": "Nota não encontrada"})
}

// DELETE: Remover nota pelo ID
func RemoverNota(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)  // converter para int
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	for i, n := range notas {
		if n.ID == id {
			notas = append(notas[:i], notas[i+1:]...)
			return c.JSON(http.StatusOK, map[string]string{"message": "Nota removida"})
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"error": "Nota não encontrada"})
}
