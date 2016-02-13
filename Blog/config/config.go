package soulconfig

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

const (
	dbConfigFile = "./Config/dbConfig.json"
)

type DbConfig struct {
	DBMS       string `json: "DBMS"`
	DbHost     string `json: "dbHost"`
	DbName     string `json: "dbName"`
	DbUsername string `json: "dbUsername"`
	DbPassword string `json: "dbPassword"`

	DbPepper string `json: "pepper"`
}

type SoulConfig struct {
	Dbc DbConfig
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
