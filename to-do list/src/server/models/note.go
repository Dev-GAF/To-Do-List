package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Nota struct {
	ID	   	 primitive.ObjectID  `json:"id" bson:"_id,omitempty"` // Os json:"..." transformam para JSON automaticamente
	Titulo 	 string  			 `json:"titulo" bson:"titulo"`
	Conteudo string  	         `json:"conteudo" bson:"conteudo"`
}