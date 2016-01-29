package mongodb

import (
	"github.com/eu271/Soulog/Blog/db"
	"github.com/eu271/Soulog/Blog/objects"
	"github.com/eu271/Soulog/TestUtil/db"
	"testing"
)

var dbConfig db.DbConfig

func TestMain(t *testing.T) {
	dbConfig = db.DbConfig{
		DBMS:       "MongoDb",
		DbHost:     "127.0.0.1",
		DbName:     "Soulog",
		DbUsername: "SoulogAdmin",
		DbPassword: "SoulogDefaultPassword",

		DbPepper: "pepperino",
	}
}

func TestMongodbQueryPost(t *testing.T) {
	var soulogDb soul.SoulogDb

	soulogDb = db.OpenDb(dbConfig)
	dbTestUtil.QueryPostTest(soulogDb, "", t)
}
