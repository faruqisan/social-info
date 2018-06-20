package twitter

import (
	"log"
	"net/http"

	"github.com/faruqisan/social-info/auth"
	"github.com/faruqisan/social-info/configs"

	"github.com/dghubble/oauth1"
	"github.com/dghubble/oauth1/twitter"
)

type Client struct {
	config *oauth1.Config
	token  *oauth1.Token
}

func NewTwitter() auth.Auth {

	cfg := configs.GetConfig()
	twCfg := cfg.TwConfig

	config := &oauth1.Config{
		ConsumerKey:    twCfg.ConsumerKey,
		ConsumerSecret: twCfg.ConsumerSecret,
		CallbackURL:    twCfg.RedirectURL,
		Endpoint:       twitter.AuthorizeEndpoint,
	}

	return &Client{
		config: config,
	}

}

func (c *Client) GetAuthorizeURL() string {

	requestToken, _, err := c.config.RequestToken()
	if err != nil {
		log.Println("Error on get temp req token : ", err)
	}

	authorizationURL, err := c.config.AuthorizationURL(requestToken)
	// handle err

	return authorizationURL.String()

}

func (c *Client) GetAccessToken(request interface{}) string {

	req, ok := request.(*http.Request)
	if !ok {
		log.Println("Request isn't *http.Request : ", ok)
	}

	requestToken, verifier, err := oauth1.ParseAuthorizationCallback(req)
	if err != nil {
		log.Println("Error on parse req token : ", err)
	}

	accessToken, accessSecret, err := c.config.AccessToken(requestToken, "", verifier)

	token := oauth1.NewToken(accessToken, accessSecret)

	c.token = token

	return token.Token

}

func (c *Client) CheckToken() bool {
	return c.token != nil
}

func (c *Client) GetAPIClient() *http.Client {
	return c.config.Client(oauth1.NoContext, c.token)
}
