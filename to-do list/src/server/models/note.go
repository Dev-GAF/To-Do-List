package models

type Nota struct {
	ID	   	 int 	 `json:"id" bson:"_id,omitempty"` // Os json:"..." transformam para JSON automaticamente
	Titulo 	 string  `json:"titulo" bson:"titulo"`
	Conteudo string  `json:"conteudo" bson:"conteudo"`
}