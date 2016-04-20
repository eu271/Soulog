package mongodb

import (
	"github.com/eu271/Soulog/Blog/objects"
	"github.com/eu271/Soulog/TestUtil/db"
	"testing"
)

var soulogDb soul.SoulogDb

func TestMain(t *testing.T) {
	soulogDb = dbTestUtil.OpenDb(OpenMongodb, t)
}

func TestMongodbQueryPost(t *testing.T) {
	dbTestUtil.QueryPostTest(soulogDb, "1", t)
	dbTestUtil.QueryPostTest(soulogDb, "1", t)

}
