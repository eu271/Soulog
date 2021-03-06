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

package soulapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/eu271/Soulog/Blog/objects"
	"log"
	"net/http"
	"strings"
	"time"
)

var soulog soul.Soulog

type json_error struct {
	Codigo  uint   `json:"codigo"`
	Mensaje string `json:"mensaje"`
}

//Returns the pettition make to the server.
func getJson(r *http.Request) *json.Decoder {
	return json.NewDecoder(r.Body)
}

func crearLlamada(nombre string, fn func(peticion *json.Decoder) string) http.HandlerFunc {

	//TODO: Check if the object passed to the API is the correct format, and
	//encode JSON/HTTP into Go Object.
	//All JSON API should be under /json/"api_call".
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" && r.Body != nil && r.ContentLength > 0 {
			http.ServeContent(w, r, nombre, time.Now(), strings.NewReader(fn(getJson(r))))
		}
	}
}

func getPost(peticion *json.Decoder) string {
	type getPostJson struct {
		Id string `json:"id"`
	}

	var p getPostJson
	var post *soul.Post
	var jsonPost string

	err := peticion.Decode(&p)
	if err != nil {
		log.Println("Se ha producido un error al decodificar una peticion de post: " + err.Error())
	}
	log.Println("Se esta pidiendo un post: " + p.Id)

	post, err = soulog.GetPost(p.Id)

	if err != nil {
		//TODO: send not found
	}

	jsonPost, err = post.PostToJson()

	if err != nil {
		//TODO: send 405 server error, conversion failed
	}

	return jsonPost
}

func getSoul(peticion *json.Decoder) string {
	log.Println("Se esta pidiendo la informacion del blog, \"Soul\"")
	return soulog.GetSoul()
}

func sendPost(peticion *json.Decoder) string {
	type sendPostJson struct {
		Titulo           string               `json:"titulo"`
		Contenido        string               `json:"contenido"`
		FechaPublicacion time.Time            `json:"fechaPublicacion"`
		Secion           autenticarSecionJson `json:"secion"`
	}
	var p sendPostJson
	err := peticion.Decode(&p)
	if err != nil {
		log.Println("Se ha producido un error al decodificar una peticion de envio de post: " + err.Error())
	}
	log.Println("Se esta enviado un post sendPost")

	b, err := autenticarSecion(p.Secion)
	if !b {
		m, _ := json.Marshal(json_error{000, err.Error()})
		return string(m)
	}
	/*
		soulog.SendPost(soul.Post{
			Id:               p.Titulo,
			Titulo:           p.Titulo,
			Contenido:        p.Contenido,
			FechaPublicacion: p.FechaPublicacion,
		})
	*/
	return "{}"
}

func deletePost(peticion *json.Decoder) string {
	type deletePostJson struct {
		Titulo string               `json:"titulo"`
		Secion autenticarSecionJson `json:"secion"`
	}
	var p deletePostJson
	err := peticion.Decode(&p)
	if err != nil {
		log.Println("Se ha producido un error al decodificar una peticion de eliminacion de post: " + err.Error())
	}
	log.Println("Se esta enviado un post deletePost")

	b, err := autenticarSecion(p.Secion)
	if !b {
		m, _ := json.Marshal(json_error{000, err.Error()})
		return string(m)
	}
	soulog.DeletePost(p.Titulo)
	return "{}"
}

func enviarPost(w http.ResponseWriter, r *http.Request) {
	var post *soul.Post
	var jsonPost string
	var err error

	if r.Method == "GET" {
		titulo := r.URL.Path[len("/post/"):]
		log.Println("Se esta pidiendo " + titulo)

		post, err = soulog.GetPost(titulo)
		if err != nil {

		}

		jsonPost, err = post.PostToJson()
		if err != nil {

		}

		http.ServeContent(w, r, "post", time.Now(), strings.NewReader(jsonPost))
	}
}
func enviarImagen(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		nombre := r.URL.Path[len("/imagen/"):]
		log.Println("Se esta pidiendo imagen " + nombre)
		http.ServeContent(w, r, nombre, time.Now(), bytes.NewReader(soulog.GetImage(nombre)))
	}
}

func imageUpload(w http.ResponseWriter, r *http.Request) {

	log.Println("Se esta subiendo una imagen")
	img, imgInfo, err := r.FormFile("imagen")
	if err != nil {
		log.Println("Error subiendo el fichero " + err.Error())
	}

	defer img.Close()

	soulog.ImagenUpload(img, imgInfo.Filename)

	fmt.Fprintf(w, "File uploaded successfully : ")
	fmt.Fprintf(w, imgInfo.Filename)
}

func AgregarFunciones(_soulog soul.Soulog) {

	soulog = _soulog

	log.Println("Setting up AJAX server API.")

	http.HandleFunc("/getPost", crearLlamada("getPost", getPost))
	http.HandleFunc("/getSoul", crearLlamada("getSoul", getSoul))

	http.HandleFunc("/getSecion", crearLlamada("getSecion", getSecion))
	http.HandleFunc("/sendPost", crearLlamada("sendPost", sendPost))
	http.HandleFunc("/deletePost", crearLlamada("deletePost", deletePost))

	http.HandleFunc("/post/", enviarPost)
	http.HandleFunc("/imagen/", enviarImagen)

	http.HandleFunc("/sendImg", imageUpload)

}
