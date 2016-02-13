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
	"github.com/eu271/Soulog/Blog/config"
	"github.com/eu271/Soulog/Blog/objects"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

const (
	DBMS = "MongoDb"
)

type mongoDb struct {
	host     string
	name     string
	user     string
	password string

	session_soulog *mgo.Session
	posts          *mgo.Collection
	images         *mgo.Collection
	users          *mgo.Collection
}

func OpenMongodb(dbc soulconfig.DbConfig) soul.SoulogDb {
	var mango mongoDb
	var err error

	mango.host = dbc.DbHost
	mango.name = dbc.DbName
	mango.user = dbc.DbUsername
	mango.password = dbc.DbPassword

	mango.session_soulog, err = mgo.Dial(mango.host)

	if err != nil {
		log.Println("Conecting with the database error. Database configuration error or deplyment with mongo unrechable: " + err.Error())
		panic("Conection open error: " + err.Error())
	}

	mango.posts = mango.session_soulog.DB(mango.name).C("Posts")
	mango.users = mango.session_soulog.DB(mango.name).C("Users")
	mango.images = mango.session_soulog.DB(mango.name).C("Images")

	return mango
}

func (mango mongoDb) QueryPostNum() uint64 {
	count, err := mango.posts.Count()
	if err != nil {
		log.Println("Error llamada GetCantidad a mongoDb. " + err.Error())
	}
	return uint64(count)
}

func (mango mongoDb) QueryPost(id string) (string, error) {
	var p []byte
	var _p soul.Post

	err := mango.posts.Find(bson.M{"id": id}).One(&_p)
	p, err = json.Marshal(_p)
	return string(p), err

}

func (mango mongoDb) GetPosts(cantidad uint64) string {
	var p string
	var _p soul.Post
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
func (mango mongoDb) InsertPost(post *soul.Post) error {
	return mango.posts.Insert(*post)
}

func (mango mongoDb) DeletePost(id string) error {
	_, err := mango.posts.RemoveAll(bson.M{"id": id})
	return err
}

func (mango mongoDb) InsertImage(image []byte, name string) error {
	return mango.images.Insert(bson.M{"id": name, "imagen": image})
}

func (mango mongoDb) QueryImage(name string) []byte {
	type img struct {
		Id     string
		Imagen string
	}
	var _p img

	log.Println(name)

	err := mango.images.Find(bson.M{"id": name}).One(&_p)
	if err != nil {
		log.Println("Error reciviendo la imagen de la base de datos: " + err.Error())
	}

	log.Println("La imagen " + _p.Id + " ha sido recogida de la base de datos.")
	return []byte(_p.Imagen)

}

func (mango mongoDb) ValidatePassword(name, password string) (bool, error) {
	var p []byte

	mango.users.Find(bson.M{"id": name}).One(&p)
	return false, nil
}
