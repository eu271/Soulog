package soulogApi

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/eu271/Soulog/Blog"
	"github.com/eu271/Soulog/Blog/objetos"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	dummySession   = "2710"
	secionCaducada = 15 //15min
)

var soulog soulogBlog.Soulog

var seciones map[string]s_secion

type s_secion struct {
	Nombre     string `json:"nombre"`
	Secion     string `json:"secion"`
	contraseña string
	Timestamp  time.Time `json:"timestamp"`
	valida     bool
	hash       string
}

type json_error struct {
	Codigo  uint   `json: "codigo"`
	Mensaje string `json:"mensaje"`
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
	type getPostJson struct {
		Id string `json:"id"`
	}

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

func getPosts(peticion *json.Decoder) string {
	type getPostsJson struct {
		Cantidad uint64 `json:"cantidad"`
	}

	var p getPostsJson
	err := peticion.Decode(&p)
	if err != nil {
		log.Println("Se ha producido un error al decodificar una peticion de post: " + err.Error())
	}

	log.Println("Se esta pidiendo una coleccion de posts")

	return soulog.GetPosts(p.Cantidad)
}

func getSecion(peticion *json.Decoder) string {
	type getSecionJson struct {
		Nombre string `json:"nombre"`
	}
	var p getSecionJson
	err := peticion.Decode(&p)
	if err != nil {

		log.Println("Se ha producido un error al decodificar una peticion de secion: " + err.Error())
	}
	log.Println("Se esta pidiendo una nueva secion")

	if seciones == nil {
		seciones = make(map[string]s_secion)
	}

	crearNumero := func() string {
		p := sha1.Sum([]byte(time.Now().String()))
		return hex.EncodeToString(p[:])
	}

	if !soulog.ExisteUsuario(p.Nombre) {
		//return error usuario incorrecto
	}

	c := soulog.GetContraseña(p.Nombre)
	s := crearNumero()
	h := sha256.Sum256([]byte(string(p.Nombre + c + s)))

	log.Println(hex.EncodeToString(h[:]))

	seciones[p.Nombre] = s_secion{
		Nombre:     p.Nombre,
		Secion:     s,
		contraseña: c,
		Timestamp:  time.Now(),
		valida:     false,
		hash:       hex.EncodeToString(h[:]),
	}

	r_secion, err := json.Marshal(seciones[p.Nombre])

	if err != nil {
		log.Println("Error serializando la secion para enviar: " + err.Error())
	}

	return string(r_secion)
}

type autenticarSecionJson struct {
	Nombre string `json:"nombre"`
	Hash   string `json:"hash"`
}

func autenticarSecion(p autenticarSecionJson) (bool, error) {

	if seciones == nil {
		return false, errors.New("No hay seciones")
	}
	_, ok := seciones[p.Nombre]
	if !ok {
		return false, errors.New("Usuario incorrecto")
	}

	if !soulog.ExisteUsuario(p.Nombre) {
		return false, errors.New("El usuario no existe")
	}

	if time.Since(seciones[p.Nombre].Timestamp).Minutes() > secionCaducada {
		return false, errors.New("Secion caducada")
	}

	if p.Hash == seciones[p.Nombre].hash {
		log.Println("Secion del usuario " + p.Nombre + " ha sido validada")
		temp, ok := seciones[p.Nombre]
		if ok {
			temp.valida = true
			seciones[p.Nombre] = temp
		}

		return true, nil
	} else {
		return false, errors.New("Secion incorrecta" + seciones[p.Nombre].hash)
	}

}

func sendPost(peticion *json.Decoder) string {
	type sendPostJson struct {
		Titulo           string               `json: "titulo"`
		Contenido        string               `json:"contenido"`
		FechaPublicacion time.Time            `json: "fechaPublicacion"`
		Secion           autenticarSecionJson `json: "secion"`
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
	soulog.SendPost(soulObjetos.Post{
		Id:               p.Titulo,
		Titulo:           p.Titulo,
		Contenido:        p.Contenido,
		FechaPublicacion: p.FechaPublicacion,
	})
	return "{}"
}

func deletePost(peticion *json.Decoder) string {
	type deletePostJson struct {
		Titulo string               `json: "titulo"`
		Secion autenticarSecionJson `json: "secion"`
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
func AgregarFunciones() {

	soulog = soulogBlog.AbrirBlog()

	log.Println("Agregando funciones a la API")

	http.HandleFunc("/getPost", crearLlamada("getPost", getPost))
	http.HandleFunc("/getTitulo", crearLlamada("getTitulo", getTitulo))
	http.HandleFunc("/getSoul", crearLlamada("getSoul", getSoul))
	http.HandleFunc("/getPosts", crearLlamada("getPosts", getPosts))

	http.HandleFunc("/getSecion", crearLlamada("getSecion", getSecion))
	http.HandleFunc("/sendPost", crearLlamada("sendPost", sendPost))
	http.HandleFunc("/deletePost", crearLlamada("deletePost", deletePost))
}
