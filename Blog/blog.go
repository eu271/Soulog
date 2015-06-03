package soulogBlog

import (
	"encoding/json"
	"github.com/eu271/Soulog/Blog/db"
	"io/ioutil"
	"log"
)

type blog struct {
	Id     string
	Titulo string
	Autor  string
	Posts  uint64

	Nombre_db     string
	Usuario_db    string
	Contraseña_db string

	soulogDb soulogDb.SoulogDb
}

func AbrirBlog() Soulog {
	var b blog

	log.Println("Cargando la configuracion del blog.")

	blogConfigString, err := ioutil.ReadFile("config.json")
	err = json.Unmarshal(blogConfigString, &b)
	if err != nil {
		log.Println("Error al abrir la configuracion del blog. " + err.Error())
	}
	return b
}

type Post struct {
	Id        string
	Titulo    string
	Contenido string
}

type Soulog interface {
	GetTitulo() string
	GetAutor() string

	GetPost(id string) string

	GetSoul() string
}

func (b blog) GetPost(id string) string {
	return "{\"titulo\":\"soy un titulo bonito\"}"
}

func (b blog) GetTitulo() string {
	return b.Titulo
}

func (b blog) GetAutor() string {
	return b.Autor
}

func (b blog) GetSoul() string {
	type soul struct {
		Id     string `json:"id"`
		Titulo string `json:"titulo"`
		Autor  string `json:"autor"`
		Posts  uint64 `json:"posts"`
	}
	d, err := json.Marshal(soul{
		b.Id,
		b.Titulo,
		b.Autor,
		b.Posts,
	})
	if err != nil {
		log.Println("Error al serializar Soul, datos del blog para enviarlos al cliente.")
	}
	return string(d)
}
