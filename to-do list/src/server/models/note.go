package models

type Nota struct {
	ID	   	 int 	 `json:"id"` // Os json:"..." servem para quando eu transformar para JSON automaticamente
	Titulo 	 string  `json:"titulo"`
	Conteudo string  `json:"conteudo"`
}