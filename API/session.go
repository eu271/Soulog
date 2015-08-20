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

	This file contains the logic to validate users and generate tokes for some of
	the API calls. At the moment should be sufficient.

	TODO: This should use a database like Redis. Where non-persistan non-critical
	data could be save in-memory.
*/

package soulogApi

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"log"
	"time"
)

const (
	dummySession   = "2710"
	secionCaducada = 15 //15min
)

var seciones map[string]s_secion

type s_secion struct {
	Nombre     string `json:"nombre"`
	Secion     string `json:"secion"`
	contraseña string
	Timestamp  time.Time `json:"timestamp"`
	valida     bool
	hash       string
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
		//TODO: Should be cryptographically secure random number.
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
