/*
	Copyright (c) 2015 Eugenio Ochoa

	Permission is hereby granted, free of charge, to any person obtaining a copy
	of this software and associated documentation files (the "Software"), to deal
	in the Software without restriction, including without limitation the rights
	to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
	copies of the Software, and to permit persons to whom the Software is
	furnished to do so, subject to the following conditions:

	The above copyright notice and this permission notice shall be included in all
	copies or substantial portions of the Software.

	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
	IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
	FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
	AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
	LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
	OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
	SOFTWARE.
*/

package soulogBlog

import (
	"bytes"
	"encoding/json"
	"github.com/eu271/Soulog/Blog/db"
	"github.com/eu271/Soulog/Blog/objects"
	"io"
	"log"
)

type blog struct {
	Id     string
	Titulo string
	Posts  uint64

	Autor      string
	Contraseña string
	salt       string

	soulogDb soulObjects.SoulogDb
}

func AbrirBlog() Soulog {
	var b blog

	b.Titulo = ""
	b.Posts = 5
	b.Autor = "Eugenio"
	b.Contraseña = "qwerty"
	b.salt = "ad"


	b.soulogDb = db.AbrirDb()

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

	SendPost(post soulObjects.Post) error
	DeletePost(id string) error

	GetImagen(nombre string) []byte
	ImagenUpload(imagen io.Reader, nombre string)
}

func (b blog) GetPost(id string) string {
	return b.soulogDb.GetPost(id)
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

func (b blog) SendPost(post soulObjects.Post) error {
	return b.soulogDb.SendPost(post)
}

func (b blog) DeletePost(id string) error {
	return b.soulogDb.DeletePost(id)
}

func (b blog) ImagenUpload(imagen io.Reader, nombre string) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(imagen)
	err := b.soulogDb.InsertarImagen(buf.Bytes(), nombre)
	if err != nil {
		log.Println("Error subiendo imagen " + err.Error())
	}
}

func (b blog) GetImagen(nombre string) []byte {
	return b.soulogDb.GetImagen(nombre)
}
