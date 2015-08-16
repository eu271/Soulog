package db

import (
	"github.com/eu271/Soulog/Blog/db/mongodb"
	"github.com/eu271/Soulog/Blog/objects"
	"encoding/json"
	"io/ioutil"
	"log"
)

const (
	DBMS_mongodb = "MongoDb"
)

type dbConfig struct {
	DBMS string `json: "DBMS"`
  DbHost string `json: "dbHost"`
  DbName string `json: "dbName"`
  DbUsername string`json: "dbUsername"`
  DbPassword string `json: "dbPassword"`

  DbPepper string `json: "pepper"`
}

func AbrirDb() soulObjects.SoulogDb {

	var dbc dbConfig

	log.Println("Loading DB configuration.")
	blogConfigString, err := ioutil.ReadFile("dbConfig.json")
	err = json.Unmarshal(blogConfigString, &dbc)
	if err != nil {
		log.Println("Error decoding dbCOnfig.json. " + err.Error())
	}

	switch dbc.DBMS {
	case DBMS_mongodb:
		return AbrirMongo(dbc.DbHost, dbc.DbName, dbc.DbUsername, dbc.DbPassword)

	}

	return AbrirMongo(dbc.DbHost, dbc.DbName, dbc.DbUsername, dbc.DbPassword)
}

func AbrirMongo(host, dbname, username, password string) soulObjects.SoulogDb {
	return mongodb.AbrirDb(host, dbname, username, password)
}
