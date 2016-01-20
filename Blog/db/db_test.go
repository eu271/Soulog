package db

import (
	"github.com/eu271/Soulog/Blog/objects"
	"testing"
)

func TestMain(t *testing.T) {

}

func queryPost(db soulObjects.SoulogDb, id string, t *testing.T) {
	t.Fatal("Fail, test not written.")
}

func TestMongodbQueryPost(t *testing.T) {
	var soulogDb soulObjects.SoulogDb
	soulogDb = OpenDb()
	queryPost(soulogDb, "", t)
}
