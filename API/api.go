package soulogApi

import (
	"encoding/json"
	"github.com/eu271/Soulog/Blog"
	"log"
	"net/http"
	"strings"
	"time"
)

var soulog soulogBlog.Soulog

type getPostJson struct {
	Id string `json:"id"`
}

//Devuelve la peticion json que se ha echo al servidor.
func getJson(r *http.Request) *json.Decoder {
	return json.NewDecoder(r.Body)
}

func crearLlamada(nombre string, fn func(peticion *json.Decoder) string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" && r.Body != nil && r.ContentLength > 0 {
			http.ServeContent(w, r, nombre, time.Now(), strings.NewReader(fn(getJson(r))))
		}
	}
}

func getPost(peticion *json.Decoder) string {
	var p getPostJson
	err := peticion.Decode(&p)
	if err != nil {
		log.Println("Se ha producido un error al decodificar una peticion de post: " + err.Error())
	}
	log.Println("Se esta pidiendo un post: " + p.Id)

	return soulog.GetPost(p.Id)
}

func getTitulo(peticion *json.Decoder) string {
	log.Println("Se esta pidiendo el titulo del blog")
	return soulog.GetTitulo()
}

func getSoul(peticion *json.Decoder) string {
	log.Println("Se esta pidiendo la informacion del blog, \"Soul\"")
	return soulog.GetSoul()
}

func AgregarFunciones() {

	soulog = soulogBlog.AbrirBlog()

	log.Println("Agregando funciones a la API")

	http.HandleFunc("/getPost", crearLlamada("getPost", getPost))
	http.HandleFunc("/getTitulo", crearLlamada("getTitulo", getTitulo))
	http.HandleFunc("/getSoul", crearLlamada("getSoul", getSoul))
}
