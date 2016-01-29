package soulConfig

import (
	"encoding/json"
	"github.com/eu271/Soulog/Blog/db"
	"io/ioutil"
	"log"
	"os"
)

const (
	dbConfigFile = "./Config/dbConfig.json"
)

type SoulConfig struct {
	Dbc db.DbConfig
}

func OpenConfig() (SoulConfig, error) {
	var soulConfig SoulConfig

	log.SetPrefix("\x1b[32mDebug:\x1b[0m ")
	log.SetOutput(os.Stdout)

	log.Println("Loading DB configuration.")
	blogConfigString, err := ioutil.ReadFile(dbConfigFile)
	if err != nil {
		log.Println("Error opening dbConfig.json " + err.Error())
	}
	err = json.Unmarshal(blogConfigString, &soulConfig.Dbc)
	if err != nil {
		log.Println("Error decoding dbCOnfig.json. " + err.Error())

		panic("Error decoding json file")
	}

	return soulConfig, err

}
