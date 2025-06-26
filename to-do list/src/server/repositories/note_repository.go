package repositories

import (
	"context"
	// "server/database"
	"server/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var notaCollection *mongo.Collection

func SetNotaCollection(coll *mongo.Collection) {
	notaCollection = coll
}

func ListarNotas() ([]models.Nota, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := notaCollection.Find(ctx, bson.M{}) // bson.M{} = busca todos os documentos
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var result []models.Nota
	for cursor.Next(ctx) {
		var n models.Nota
		if err := cursor.Decode(&n); err != nil { // Converte o documento BSON em struct Go
			return nil, err // trata erro de decode
		}
		result = append(result, n)
	}

	if result == nil {
		result = []models.Nota{}
	}

	return result, nil
}

func CriarNota(nota models.Nota) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := notaCollection.InsertOne(ctx, nota)
	return err
}

func AtualizarNota(id primitive.ObjectID, nota models.Nota) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"titulo": nota.Titulo,
			"conteudo": nota.Conteudo,
		},
	}

	_, err := notaCollection.UpdateOne(ctx, bson.M{"_id": id}, update)
	return err
}

func RemoverNota(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := notaCollection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}