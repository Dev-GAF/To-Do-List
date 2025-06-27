package controllers

import (
	"net/http"
	// "time"
	// "strconv"

	"github.com/labstack/echo/v4"

	"server/models"
	"server/repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GET: Obter nota por ID
func ObterNotaPorID(c echo.Context) error {
	idStr := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "ID inválido"})
	}

	nota, err := repositories.ObterNotaPorID(objectID)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Nota não encontrada"})
	}

	return c.JSON(http.StatusOK, nota)
}


// GET: Obter todos as notas
func ObterNotas(c echo.Context) error {
	notas, err := repositories.ObterNotas()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Erro ao buscar"})
	}

	return c.JSON(http.StatusOK, notas)
}

// POST: Criar nova nota
func CriarNota(c echo.Context) error {
	var nota models.Nota
	if err := c.Bind(&nota); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Dados inválidos"})
	}

	nota.ID = primitive.NewObjectID() // gera novo ObjectID para o MongoDB

	err := repositories.CriarNota(nota)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Erro ao salvar"})
	}

	return c.JSON(http.StatusCreated, nota)
}

// PUT: Atualizar nota pelo ID
func AtualizarNota(c echo.Context) error {
	idStr := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "ID inválido"})
	}

	var notaAtualizada models.Nota
	if err := c.Bind(&notaAtualizada); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if notaAtualizada.ID != primitive.NilObjectID && notaAtualizada.ID != objectID {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "ID do corpo diferente do ID da URL"})
	}
	notaAtualizada.ID = objectID

	err = repositories.AtualizarNota(objectID, notaAtualizada)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Erro ao atualizar nota"})
	}

	return c.JSON(http.StatusOK, notaAtualizada)
}

// DELETE: Remover nota pelo ID
func RemoverNota(c echo.Context) error {
	idStr := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(idStr) // strconv.Atoi(idStr)  // converter para int
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "ID inválido"})
	}

	err = repositories.RemoverNota(objectID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Erro ao remover nota"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Nota removida com sucesso"})
}
