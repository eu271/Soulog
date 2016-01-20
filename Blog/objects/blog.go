package soulObjects

import (
	"io"
)

type Soulog interface {
	GetPost(id string) string
	GetSoul() string

	ExisteUsuario(name string) bool
	GetContraseña(name string) string
	LoginUser(name, password string) bool

	SendPost(post Post) error
	DeletePost(id string) error

	GetImage(nombre string) []byte
	ImagenUpload(imagen io.Reader, nombre string)
}
