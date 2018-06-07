package configs

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// DatabaseConfig .
type DatabaseConfig struct {
	Host     string `json:"database-host"`
	Database string `json:"database-name"`
	Username string `json:"database-username"`
	Password string `json:"database-password"`
}

func loadDatabaseConfig() (databaseConfig DatabaseConfig, err error) {

	env, err := ioutil.ReadFile("../env.json")
	if err != nil {
		log.Fatalln("Unable to read env file:", err)
		return
	}

	err = json.Unmarshal(env, &databaseConfig)
	if err != nil {
		log.Fatalln("Unable to unmarshal:", err)
	}

	return
}
