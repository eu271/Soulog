package dbTestUtil

import (
	"github.com/eu271/Soulog/Blog/config"
	"github.com/eu271/Soulog/Blog/objects"
	"github.com/eu271/Soulog/TestUtil/objects"
	"testing"
	"time"
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
func QueryPostTest(db soul.SoulogDb, t *testing.T) {

	var postToDb *soul.Post

	postToDb = objectsTestUtil.NewTestPost()

	InsertPostTest(db, postToDb, t)
}

//Test for the function InsertPost in the interface SoulogDb. Used in the
//DBMS tests.
func InsertPostTest(db soul.SoulogDb, post *soul.Post, t *testing.T) {

	var postFromDb *soul.Post

	err := db.InsertPost(post)
	if err != nil {
		t.Error("Error inserting post: " + err.Error())
	}

	postFromDb, err = db.QueryPost(post.Id)
	if err != nil {
		t.Error("Error getting post: " + err.Error())
	}

	if postFromDb == nil {
		t.Error("Post form db is nil")
	}

	if post.Id != postFromDb.Id {
		t.Error("The inserted post is not the same as the one retrive from db. Initial post id=" + post.Id + " Recive post id=" + postFromDb.Id)
	}
}

func QueryPostBetweenDates(db soul.SoulogDb, from, to time.Time, t *testing.T) {
	var post1, post2, post3 *soul.Post

}

func OpenDb(openDb func(dbConfig soulconfig.DbConfig) soul.SoulogDb, t *testing.T) soul.SoulogDb {
	return openDb(dbConfig)
}
