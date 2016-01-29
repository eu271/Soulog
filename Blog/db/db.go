package db

import (
	"github.com/eu271/Soulog/Blog/db/mongodb"
	"github.com/eu271/Soulog/Blog/objects"
	"log"
	//"errors"
)

const (
	DBMS_mongodb = "MongoDb"
)

type DbConfig struct {
	DBMS       string `json: "DBMS"`
	DbHost     string `json: "dbHost"`
	DbName     string `json: "dbName"`
	DbUsername string `json: "dbUsername"`
	DbPassword string `json: "dbPassword"`

	DbPepper string `json: "pepper"`
}

func OpenDb(dbc DbConfig) soul.SoulogDb {

	switch dbc.DBMS {
	case DBMS_mongodb:
		log.Println("Opening MongoDb as the main database.")
		return mongodb.OpenMongodb(dbc.DbHost, dbc.DbName, dbc.DbUsername, dbc.DbPassword)

	}

	return nil //errors.New("DBMS not selected or invalid.")
}
