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
	"io/ioutil"
	"encoding/json"
	//"time"
	//"github.com/eu271/Soulog/Blog/objetos"

	"github.com/eu271/Soulog/API"
)

const (
	Index      = "./cliente/blog.html"
)

type fileToServeStruct struct {
	ServeFiles []struct {
		FilePath string `json: "filePath"`
		FilePathServe string `json: "filePathServe"`
		ContentType string `json: "contentType"`
	}
}

func ServirIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, Index)
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
	http.ServeFile(w, r, Index)
}

func loadFiles() {
	var filesToServe fileToServeStruct

	filesToServeString, _ := ioutil.ReadFile("filesToServe.json")
	err := json.Unmarshal(filesToServeString, &filesToServe)
	if err != nil {
		log.Println("Error decoding the default configuration. " + err.Error())
	}

	log.Println("Adding BlogResources to the URLs available to send.")

	for _, f := range filesToServe.ServeFiles {
		registrarDireccionFichero(f.FilePathServe, f.FilePath, f.ContentType)
	}
}

func main() {



	log.SetPrefix("Debug: ")
	log.SetOutput(os.Stdout)

	loadFiles()

	soulogApi.AgregarFunciones()

	http.HandleFunc("/", ServirIndex)


	///*
	go func() {
		log.Println("Starting https login webpage...")
		err := http.ListenAndServeTLS(":8088", "claves_ssl/cert.pem", "claves_ssl/key.pem", http.HandlerFunc(ServeTLS))

		log.Println("Error in tls server: " + err.Error())
	}()
	//*/

	log.Println("Starting server...")
	err := http.ListenAndServe(":8080", nil)

	log.Println("Error in http server: " + err.Error())
}
