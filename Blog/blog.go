package soulogBlog

import (
	"encoding/json"
	"github.com/eu271/Soulog/Blog/db"
	"github.com/eu271/Soulog/Blog/objetos"
	"io/ioutil"
	"log"
)

type blog struct {
	Id     string
	Titulo string
	Posts  uint64

	Autor      string
	Contraseña string
	salt       string

	Host_db       string
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

	b.soulogDb = soulogDb.AbrirDb(b.Host_db, b.Nombre_db, b.Usuario_db, b.Contraseña_db)

	b.Posts = b.soulogDb.GetCantidad()

	log.Println(b.Posts)

	return b
}

type Soulog interface {
	GetTitulo() string
	GetAutor() string
	GetPost(id string) string
	GetPosts(cantidad uint64) string
	GetSoul() string
	ExisteUsuario(Nombre string) bool
	GetContraseña(Nombre string) string

	SendPost(post soulObjetos.Post) error
	DeletePost(id string) error
}

func (b blog) GetPost(id string) string {
	return "{\"titulo\":\"soy un titulo bonito\"}"
}

func (b blog) GetPosts(cantidad uint64) string {
	if b.Posts < cantidad {
		cantidad = b.Posts
	}

	return b.soulogDb.GetPosts(cantidad)
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

func (b blog) ExisteUsuario(Nombre string) bool {
	return Nombre == "Eugenio"
}

func (b blog) GetContraseña(Nombre string) string {
	return "65e84be33532fb784c48129675f9eff3a682b27168c0ea744b2cf58ee02337c5" //qwerty
}

func (b blog) SendPost(post soulObjetos.Post) error {
	return b.soulogDb.SendPost(post)
}

func (b blog) DeletePost(id string) error {
	return b.soulogDb.DeletePost(id)
}
