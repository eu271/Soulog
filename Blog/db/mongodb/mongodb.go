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

package mongodb

import (
	"encoding/json"
	"github.com/eu271/Soulog/Blog/objects"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type mongoDb struct {
	host       string
	nombre     string
	usuario    string
	contraseña string

	session_soulog *mgo.Session
	posts          *mgo.Collection
	imagenes       *mgo.Collection
	users          *mgo.Collection
}

func AbrirDb(host, nombre, usuario, contraseña string) soulObjects.SoulogDb {
	var mango mongoDb
	var err error

	mango.host = host
	mango.nombre = nombre
	mango.usuario = usuario
	mango.contraseña = contraseña

	log.Println("Abriendo base de datos MongoDb.")

	mango.session_soulog, err = mgo.Dial(mango.host)

	if err != nil {
		log.Println("Error en la conexion y apertura de la base de datos. " + err.Error())
	}

	mango.posts = mango.session_soulog.DB(mango.nombre).C("Posts")
	mango.users = mango.session_soulog.DB(mango.nombre).C("Users")
	mango.imagenes = mango.session_soulog.DB(mango.nombre).C("Imagenes")

	return mango
}

func (mango mongoDb) GetCantidad() uint64 {
	count, err := mango.posts.Count()
	if err != nil {
		log.Println("Error llamada GetCantidad a mongoDb. " + err.Error())
	}
	return uint64(count)
}

func (mango mongoDb) GetPost(id string) string {
	var p []byte
	var _p soulObjects.Post

	_ = mango.posts.Find(bson.M{"id": id}).One(&_p)
	p, _ = json.Marshal(_p)
	return string(p)

}

func (mango mongoDb) GetPosts(cantidad uint64) string {
	var p string
	var _p soulObjects.Post
	i := mango.posts.Find(nil).Iter()
	p = "["
	for i.Next(&_p) {
		_s, _ := json.Marshal(_p)
		p = p + string(_s) + ","
	}

	p = p[:len(p)-1]
	p = p + "]"

	return p
}
func (mango mongoDb) SendPost(post soulObjects.Post) error {
	return mango.posts.Insert(post)
}

func (mango mongoDb) DeletePost(id string) error {
	_, err := mango.posts.RemoveAll(bson.M{"id": id})
	return err
}

func (mango mongoDb) InsertarImagen(imagen []byte, nombre string) error {
	return mango.imagenes.Insert(bson.M{"id": nombre, "imagen": imagen})
}

func (mango mongoDb) GetImagen(nombre string) []byte {
	type img struct {
		Id     string
		Imagen string
	}
	var _p img

	log.Println(nombre)

	err := mango.imagenes.Find(bson.M{"id": nombre}).One(&_p)
	if err != nil {
		log.Println("Error reciviendo la imagen de la base de datos: " + err.Error())
	}

	log.Println("La imagen " + _p.Id + " ha sido recogida de la base de datos.")
	return []byte(_p.Imagen)

}
