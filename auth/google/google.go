// Sample Go code for user authorization

package google

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"google.golang.org/api/plus/v1"

	"github.com/faruqisan/social-info/database/postgresql"

	"github.com/faruqisan/social-info/auth"
	"github.com/faruqisan/social-info/database"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"
)

const missingClientSecretsMessage = `Please configure OAuth 2.0`

var config *oauth2.Config
var db database.Database

// Credentials store .json cred into struct
type Credentials struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectURL  string `json:"redirect_uris"`
}

// API struct
type API struct {
	Token *oauth2.Token
}

type SavedToken struct {
	AccessToken  string    `db:"access_token"`
	RefreshToken string    `db:"refresh_token"`
	TokenType    string    `db:"token_type"`
	Expiry       time.Time `db:"expiry"`
	Email        string    `db:"email"`
	ID           int64     `db:"id"`
}

// NewGoogleAPI return instance of Client that implemented Auth interface
func NewGoogleAPI() auth.Auth {

	b, err := ioutil.ReadFile("../google-web-api.json")
	if err != nil {
		log.Panicln("Unable to read client secret file:", err)
	}

	scopes := []string{
		youtube.YoutubeReadonlyScope,
		youtube.YoutubeUploadScope,
		plus.UserinfoEmailScope,
	}

	config, err = google.ConfigFromJSON(b, scopes...)
	if err != nil {
		log.Panicln("Unable to parse client secret file to config: ", err)
	}

	db = postgresql.NewPostgresql()

	tok, err := getTokenFromDatabase()
	if err != nil {
		log.Panicln("Unable to get token from db: ", err)
	}

	return &API{
		tok,
	}
}

// GetAccessToken return token from oauth
func (g *API) GetAccessToken(authCode interface{}) string {
	ctx := context.Background()

	// check token from database

	dbToken, err := getTokenFromDatabase()
	if err != nil {
		log.Println("Error get token from database", err)
	}

	if dbToken != nil {

		g.Token = dbToken

		now := time.Now()

		if dbToken.Expiry.Before(now) {
			g.refreshToken()
			log.Println("token refreshed")
		}

		return dbToken.AccessToken
	}

	aCode, ok := authCode.(string)
	if !ok {
		log.Println("Request isn't string : ", ok)
	}

	token, err := config.Exchange(ctx, aCode)
	if err != nil {
		log.Println("Unable to exchange tokens:", err)
	}

	log.Println("token inside get access token :", token)

	g.Token = token

	err = g.saveToken(token)
	if err != nil {
		log.Panicln("Fail to save token to database : ", err)
	}

	return token.AccessToken
}

// GetAuthorizeURL return redirect url for user to login into their google account
func (g *API) GetAuthorizeURL() string {
	return config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
}

// GetAPIClient wrap getClient
func (g *API) GetAPIClient() *http.Client {
	ctx := context.Background()

	return config.Client(ctx, g.Token)
}

// CheckToken .
func (g *API) CheckToken() bool {

	return g.Token != nil

}

func (g *API) refreshToken() {
	ctx := context.Background()

	token, err := config.Exchange(ctx, g.Token.RefreshToken)
	if err != nil {
		log.Panicln("Unable to exchange tokens: ", err)
	}

	g.Token = token
}

func (g *API) saveToken(token *oauth2.Token) (err error) {

	gc := g.GetAPIClient()

	plusService, err := plus.New(gc)
	if err != nil {
		log.Fatalf("Unable to init G+ Client %v", err)
	}

	person, err := plusService.People.Get("me").Do()
	if err != nil {
		log.Fatalf("Unable to init G+ Client %v", err)
	}

	// Store access token to db

	db, err := db.GetDatabase()
	if err != nil {
		log.Panicln(err)
	}

	stmt, err := db.Prepare(db.RebindMaster(insertAccessToken))
	if err != nil {
		log.Panicln(err)
	}

	var email *plus.PersonEmails

	for _, e := range person.Emails {
		if e.Type == "account" {
			email = e
		}
	}

	_, err = stmt.Exec(token.AccessToken, token.RefreshToken, token.TokenType, token.Expiry, email.Value)
	if err != nil {
		log.Panicln(err)
	}

	return
}

func getTokenFromDatabase() (*oauth2.Token, error) {

	var savedToken SavedToken

	db, err := db.GetDatabase()
	if err != nil {
		log.Fatalln(err)
	}

	row := db.QueryRowx(getAccessToken)
	err = row.StructScan(&savedToken)
	if err != nil {
		log.Println("no result", err)
		return nil, err
	}

	return &oauth2.Token{
		AccessToken:  savedToken.AccessToken,
		TokenType:    savedToken.TokenType,
		RefreshToken: savedToken.RefreshToken,
		Expiry:       savedToken.Expiry,
	}, nil

}
