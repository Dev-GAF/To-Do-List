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

var checklistCollection *mongo.Collection

func SetChecklistCollection(coll *mongo.Collection) {
	checklistCollection = coll
}

func ListarChecklists() ([]models.Checklist, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := checklistCollection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var result []models.Checklist
	for cursor.Next(ctx) {
		var c models.Checklist
		if err := cursor.Decode(&c); err != nil {
			return nil, err
		}
		result = append(result, c)
	}

	return result, nil
}

func CriarChecklist(checklist models.Checklist) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := checklistCollection.InsertOne(ctx, checklist)
	return err
}

func AtualizarChecklist(id primitive.ObjectID, checklist models.Checklist) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for i := range checklist.Items {
		if checklist.Items[i].ID.IsZero() {
			newID := primitive.NewObjectID()
			checklist.Items[i].ID = newID
		}
	}

	update := bson.M{
		"$set": bson.M{
			"titulo": checklist.Titulo,
			"items":  checklist.Items,
		},
	}

	_, err := checklistCollection.UpdateOne(ctx, bson.M{"_id": id}, update)
	return err
}

func RemoverChecklist(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := checklistCollection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
