package dbTestUtil

import (
	"github.com/eu271/Soulog/Blog/objects"
	"github.com/eu271/Soulog/TestUtil/objects"
	"testing"
)

//Test for the function QueryPost in the interface SoulogDb. Used in the
//DBMS tests.
func QueryPostTest(db soul.SoulogDb, id string, t *testing.T) {

	var p soul.Post
	var postJson string

	p = objectsTestUtil.NewTestPost()

	err := db.InsertPost(p)
	if err != nil {
		t.Error("Error inserting post: " + err.Error())
	}

	postJson, err = db.QueryPost(p.Id)
	if err != nil {
		t.Error("Error getting post: " + err.Error())
	}

	pReturned, _ := soul.NewPostBuilder().Json(postJson)

	if pReturned.Id != p.Id {
		t.Error("The inserted post is not he same as the one retrive form the db. Initial post id=" + p.Id + " Recive post id=" + pReturned.Id)
	}
}
