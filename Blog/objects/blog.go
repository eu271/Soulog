package soul

import (
	"io"
)

type Soulog interface {
	GetPost(id string) (*Post, error)
	SavePost(post Post) error
	DeletePost(id string) error

	GetSoul() string

	ExisteUsuario(name string) bool
	GetContraseña(name string) string
	LoginUser(name, password string) bool

	GetImage(nombre string) []byte
	ImagenUpload(imagen io.Reader, nombre string)
}
