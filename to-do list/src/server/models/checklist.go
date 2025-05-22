package models

type Checklist struct {
	ID     string           `json:"id"`
	Titulo string           `json:"titulo"`
	Items  []ChecklistItem  `json:"items"`
}

type ChecklistItem struct {
	ID     string  `json:"id"`
	Titulo string  `json:"titulo"`
	Feito  bool    `json:"feito"`
}


