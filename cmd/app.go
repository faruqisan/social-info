package main

import (
	"log"
	"net/http"
	"time"

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

	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()
	go func() {
		for t := range ticker.C {
			log.Println(t)
		}
	}()

	log.Printf("****Server Running on Port %s****", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
