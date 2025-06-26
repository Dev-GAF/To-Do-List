package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Checklist struct {
	ID     primitive.ObjectID  `json:"id" bson:"_id,omitempty"`
	Titulo string          	   `json:"titulo" bson:"titulo"`
	Items  []ChecklistItem 	   `json:"items" bson:"items"`
}

type ChecklistItem struct {
	ID     primitive.ObjectID  `json:"id" bson:"_id,omitempty"`
	Titulo string 			   `json:"titulo" bson:"titulo"`
	Feito  bool   			   `json:"feito" bson:"feito"`
}
