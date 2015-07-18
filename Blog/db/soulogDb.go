package soulogDb

import (
	"encoding/json"
	"github.com/eu271/Soulog/Blog/objetos"
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
}

func AbrirDb(host, nombre, usuario, contraseña string) SoulogDb {
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

	return mango
}

type SoulogDb interface {
	GetPost(id string) string
	GetPosts(cantidad uint64) string
	GetCantidad() uint64

	SendPost(post soulObjetos.Post) error
	DeletePost(id string) error
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
	var _p soulObjetos.Post

	_ = mango.posts.Find(bson.M{"id": id}).One(&_p)
	p, _ = json.Marshal(_p)
	return string(p)

}

func (mango mongoDb) GetPosts(cantidad uint64) string {
	var p string
	var _p soulObjetos.Post
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
func (mango mongoDb) SendPost(post soulObjetos.Post) error {
	return mango.posts.Insert(post)
}

func (mango mongoDb) DeletePost(id string) error {
	_, err := mango.posts.RemoveAll(bson.M{"id": id})
	return err
}
