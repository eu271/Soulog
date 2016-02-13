package dbTestUtil

import (
	"github.com/eu271/Soulog/Blog/config"
	"github.com/eu271/Soulog/Blog/objects"
	"github.com/eu271/Soulog/TestUtil/objects"
	"testing"
)

var dbConfig = soulconfig.DbConfig{
	DBMS:       "MongoDb",
	DbHost:     "127.0.0.1:27017",
	DbName:     "Soulog",
	DbUsername: "SoulogAdmin",
	DbPassword: "SoulogDefault",

	DbPepper: "Mi43MTgyODE4Mjg0NTkwNDUyMzUzNjAyODc0NzEzNTI2NjI0OTc=",
}

//Test for the function QueryPost in the interface SoulogDb. Used in the
//DBMS tests.
func QueryPostTest(db soul.SoulogDb, id string, t *testing.T) {

	var postToDb *soul.Post
	var postFromDb *soul.Post
	var postJson string

	postToDb = objectsTestUtil.NewTestPost()

	err := db.InsertPost(postToDb)
	if err != nil {
		t.Error("Error inserting post: " + err.Error())
	}

	postJson, err = db.QueryPost(postToDb.Id)
	if err != nil {
		t.Error("Error getting post: " + err.Error())
	}

	postFromDb, err = soul.NewPostBuilder().Json(postJson)

	if err != nil {
		t.Error(err.Error())
	}

	if postFromDb == nil {
		t.Error("Post from db is nil")
	}

	if postFromDb.Id != postToDb.Id {
		t.Error("The inserted post is not he same as the one retrive form the db. Initial post id=" + postToDb.Id + " Recive post id=" + postFromDb.Id)
	}
}

func OpenDb(openDb func(dbConfig soulconfig.DbConfig) soul.SoulogDb, t *testing.T) soul.SoulogDb {
	return openDb(dbConfig)
}
