package handlers

import (
	"log"
	"net/http"

	"github.com/faruqisan/social-info/auth"
	twitterAuth "github.com/faruqisan/social-info/auth/twitter"
	"github.com/faruqisan/social-info/src/twitter"
	"github.com/gin-gonic/gin"
)

type TwitterHandler struct {
	twitterAuth auth.Auth
}

func NewTwitterHandler() TwitterHandler {

	twitterAuth := twitterAuth.NewTwitter()

	return TwitterHandler{
		twitterAuth: twitterAuth,
	}
}

func (t *TwitterHandler) Authorize(c *gin.Context) {

	if !t.twitterAuth.CheckToken() {
		c.Redirect(http.StatusFound, t.twitterAuth.GetAuthorizeURL())
	}

}

func (t *TwitterHandler) HandleCallback(c *gin.Context) {

	t.twitterAuth.GetAccessToken(c.Request)

	c.JSON(200, t.twitterAuth.CheckToken())

}

func (t *TwitterHandler) SendTweet(c *gin.Context) {

	if !t.twitterAuth.CheckToken() {
		t.Authorize(c)
	}

	client := twitter.NewTwitterClient(t.twitterAuth)

	text := c.PostForm("tweet")

	tweet, resp, err := client.SendTweet(text)

	if err != nil {
		log.Println(err)
	}

	if resp.StatusCode == 200 {

		c.JSON(200, tweet)

	}

}
