package models

type Checklist struct {
	ID     int 	 			`json:"id"`
	Titulo string           `json:"titulo"`
	Items  []ChecklistItem  `json:"items"`
}

type ChecklistItem struct {
	ID     int 	   `json:"id"`
	Titulo string  `json:"titulo"`
	Feito  bool    `json:"feito"`
}


