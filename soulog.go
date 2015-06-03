package main

import (
	"log"
	"net/http"
	"os"

	"github.com/eu271/Soulog/API"
)

const (
	ConfigFile = "./config.json"
	Index      = "./cliente/blog.html"
)

func ServirIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, Index)
}

func registrarDireccionFichero(patron string, fichero string, tipo string) {
	http.HandleFunc(patron, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("type", tipo)
		http.ServeFile(w, r, fichero)
	})
}

func main() {

	log.SetPrefix("Debug: ")
	log.SetOutput(os.Stdout)

	soulogApi.AgregarFunciones()

	http.HandleFunc("/", ServirIndex)

	log.Println("Registrando las direcciones del blog.")
	registrarDireccionFichero("/blog.css", "./cliente/blog.css", "text/css")
	registrarDireccionFichero("/blog.js", "./cliente/blog.js", "application/javascript")

	log.Println("Iniciando el servidor.")
	http.ListenAndServe(":8080", nil)
}
