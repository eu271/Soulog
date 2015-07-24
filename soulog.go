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

package main

import (
	"log"
	"net/http"
	"os"
	
	"time"
	"github.com/eu271/Soulog/Blog/objetos"
	
	"github.com/eu271/Soulog/API"
)

const (
	ConfigFile = "./config.json"
	Index      = "./cliente/blog.html"
)

func ServirIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, Index)
}

func ServirIndex1(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, ConfigFile)
}

func registrarDireccionFichero(patron string, fichero string, tipo string) {
	http.HandleFunc(patron, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("type", tipo)
		http.ServeFile(w, r, fichero)
	})
}
	
func ServeTLS(w http.ResponseWriter, r *http.Request) {
// Generate the developer certificates:
//openssl req -newkey rsa:2048 -new -nodes -x509 -days 3650 -keyout key.pem -out cert.pem

	log.Println("Asking for loging webpage.")
	http.ServeFile(w, r, ConfigFile)
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
	registrarDireccionFichero("/javascript_global/commonmark.js", "./cliente/javascript_global/commonmark.js", "application/javascript")
	registrarDireccionFichero("/javascript_global/jquery.mustache.js", "./cliente/javascript_global/jquery.mustache.js", "application/javascript")



	log.Println("Adding AdminResourses to the URLs available to send.")
	registrarDireccionFichero("/admin", "./cliente/soul/soul.html", "text/html")
	registrarDireccionFichero("/soul.css", "./cliente/soul/soul.css", "text/css")
	registrarDireccionFichero("/soul.js", "./cliente/soul/soul.js", "application/javascript")
	
	log.Println("Adding Templates to the URLs available to send.")
	registrarDireccionFichero("/templates/post.html.mustache", "./cliente/templates/post.html.mustache", "text/html")


	
	///*
	go func(){
		log.Println("Starting https login webpage...")
		err := http.ListenAndServeTLS(":8088", "claves_ssl/cert.pem", "claves_ssl/key.pem", http.HandlerFunc(ServeTLS))
		
		log.Println("Error in tls server: " + err.Error())
	}()
	//*/
	
	log.Println("Starting server...")
	err := http.ListenAndServe(":8080", nil)
	
	log.Println("Error in http server: " + err.Error())
}
