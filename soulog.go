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

	log.Println("Adding BlogResources to the URLs available to send.")
	registrarDireccionFichero("/blog.css", "./cliente/blog.css", "text/css")
	registrarDireccionFichero("/blog.js", "./cliente/blog.js", "application/javascript")
	registrarDireccionFichero("/javascript_global/sha256.js", "./cliente/javascript_global/sha256.js", "application/javascript")
	registrarDireccionFichero("/javascript_global/variables.js", "./cliente/javascript_global/variables.js", "application/javascript")
	registrarDireccionFichero("/javascript_global/mmd.js", "./cliente/javascript_global/mmd.js", "application/javascript")

	log.Println("Adding AdminResourses to the URLs available to send.")
	registrarDireccionFichero("/admin", "./cliente/soul/soul.html", "text/html")
	registrarDireccionFichero("/soul.css", "./cliente/soul/soul.css", "text/css")
	registrarDireccionFichero("/soul.js", "./cliente/soul/soul.js", "application/javascript")
	
	log.Println("Adding Templates to the URLs available to send.")
	registrarDireccionFichero("/post.html.mustache", "./cliente/templates/post.html.mustache", "text/html")

	log.Println("Starting server...")
	/*
			go func(){
				http.ListenAndServeTLS(":8088", )
		}
		+*/
	http.ListenAndServe(":8080", nil)
}
