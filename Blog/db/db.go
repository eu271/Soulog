package db

import (
	"encoding/json"
	"github.com/eu271/Soulog/Blog/db/mongodb"
	"github.com/eu271/Soulog/Blog/objects"
	"io/ioutil"
	"log"
	//"errors"
)

const (
	DBMS_mongodb = "MongoDb"
)

type dbConfig struct {
	DBMS       string `json: "DBMS"`
	DbHost     string `json: "dbHost"`
	DbName     string `json: "dbName"`
	DbUsername string `json: "dbUsername"`
	DbPassword string `json: "dbPassword"`

	DbPepper string `json: "pepper"`
}

func OpenDb() soulObjects.SoulogDb {

	var dbc dbConfig

	log.Println("Loading DB configuration.")
	blogConfigString, err := ioutil.ReadFile("dbConfig.json")
	err = json.Unmarshal(blogConfigString, &dbc)
	if err != nil {
		log.Println("Error decoding dbCOnfig.json. " + err.Error())
	}

	switch dbc.DBMS {
	case DBMS_mongodb:
		log.Println("Opening MongoDb as the main database.")
		return mongodb.OpenMongodb(dbc.DbHost, dbc.DbName, dbc.DbUsername, dbc.DbPassword)

	}

	return nil //errors.New("DBMS not selected or invalid.")
}
