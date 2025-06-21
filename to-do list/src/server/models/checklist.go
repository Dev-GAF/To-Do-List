package models

type Checklist struct {
	ID     int 	 			`json:"id" "_id,omitempty"`
	Titulo string           `json:"titulo" bson:"titulo"`
	Items  []ChecklistItem  `json:"items" bson:"items"`
}

type ChecklistItem struct {
	ID     int 	   `json:"id" bson:"id"`
	Titulo string  `json:"titulo" bson:"tarefa"`
	Feito  bool    `json:"feito" bson:"feito"`
}


