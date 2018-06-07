package auth

import (
	"net/http"
)

// Auth interface is a contract for implementors
type Auth interface {
	GetAuthorizeURL() string
	GetAccessToken(authCode string) string
	CheckToken() bool
	GetAPIClient() *http.Client
}
