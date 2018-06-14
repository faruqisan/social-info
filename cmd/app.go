package main

import (
	"log"

	"github.com/faruqisan/social-info/configs"

	"github.com/faruqisan/social-info/services/api"
)

const port = ":8080"

func main() {

	err := configs.InitConfig()
	if err != nil {
		log.Fatalln("error init config")
	}

	api.RegisterAPI()

}
