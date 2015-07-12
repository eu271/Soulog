package soulObjetos

import (
	"time"
)

type Post struct {
	Id               string    `json:"id"`
	Titulo           string    `json:"titulo"`
	Contenido        string    `json:"contenido"`
	FechaPublicacion time.Time `json:"fechaPublicacion"`
}
