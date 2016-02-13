package db

import (
	"github.com/eu271/Soulog/Blog/config"
	"github.com/eu271/Soulog/Blog/db/mongodb"
	"github.com/eu271/Soulog/Blog/objects"
	"log"
	//"errors"
)

const (
	DBMS_mongodb = "MongoDb"
)

func OpenDb(dbc soulconfig.DbConfig) soul.SoulogDb {

	switch dbc.DBMS {
	case DBMS_mongodb:
		log.Println("Opening MongoDb as the main database.")
		return mongodb.OpenMongodb(dbc)

	}

	return nil //errors.New("DBMS not selected or invalid.")
}
