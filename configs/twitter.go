package configs

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// DatabaseConfig .
type TwitterConfig struct {
	ConsumerKey    string `json:"twitter-consumer-key"`
	ConsumerSecret string `json:"twitter-consumer-secret"`
	RedirectURL    string `json:"twitter-redirect-url"`
}

func loadTwitterConfig() (twitterConfig TwitterConfig, err error) {

	env, err := ioutil.ReadFile("../env.json")
	if err != nil {
		log.Fatalln("Unable to read env file:", err)
		return
	}

	err = json.Unmarshal(env, &twitterConfig)
	if err != nil {
		log.Fatalln("Unable to unmarshal:", err)
	}

	return
}
