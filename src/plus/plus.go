package plus

import (
	"log"
	"net/http"

	"google.golang.org/api/plus/v1"
)

type Client struct {
	service *plus.Service
}

func NewPlusCliient(googleClient *http.Client) Client {

	service, err := plus.New(googleClient)
	if err != nil {
		log.Fatalf("Unable to init G+ Client %v", err)
	}

	return Client{service}
}

func (c *Client) GetPeople() *plus.Person {

	person, err := c.service.People.Get("me").Do()
	if err != nil {
		log.Fatalf("Unable to init G+ Client %v", err)
	}
	return person

}
