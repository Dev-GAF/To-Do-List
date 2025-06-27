package repositories

import (
	"context"
	// "server/database"
	"server/models"
	"time"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var checklistCollection *mongo.Collection

func SetChecklistCollection(coll *mongo.Collection) {
	checklistCollection = coll
}

func ListarChecklistsPorId(id primitive.ObjectID) (models.Checklist, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var checklist models.Checklist
	err := checklistCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&checklist)
	
	return checklist, err
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

func AtualizarChecklist(id primitive.ObjectID, checklist models.Checklist) (models.Checklist, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var checklistExistente models.Checklist
	err := checklistCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&checklistExistente)
	if err != nil {
		return models.Checklist{}, err
	}

	existingItemIDs := make(map[primitive.ObjectID]bool)
	existingTitles := make(map[string]primitive.ObjectID)

	for _, item := range checklistExistente.Items {
		if item.ID != nil {
			existingItemIDs[*item.ID] = true
		}
		existingTitles[item.Titulo] = *item.ID 
	}

	var novosItems []models.ChecklistItem
	for _, item := range checklist.Items {
		if item.ID != nil && !item.ID.IsZero() {
			if !existingItemIDs[*item.ID] {
				return models.Checklist{}, fmt.Errorf("ID de item inválido: %s", item.ID.Hex())
			}
			novosItems = append(novosItems, item)
		} else {
			if _, exists := existingTitles[item.Titulo]; exists {
				return models.Checklist{}, fmt.Errorf("Item \"%s\" já existe e precisa de ID no corpo da requisição", item.Titulo)
			}
			newID := primitive.NewObjectID()
			item.ID = &newID
			novosItems = append(novosItems, item)
		}
	}

	update := bson.M{
		"$set": bson.M{
			"titulo": checklist.Titulo,
			"items":  novosItems,
		},
	}

	_, err = checklistCollection.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		return models.Checklist{}, err
	}

	checklist.ID = id
	checklist.Items = novosItems
	return checklist, nil
}


func RemoverChecklist(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := checklistCollection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
