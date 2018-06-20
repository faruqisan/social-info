package handlers

import (
	"github.com/faruqisan/social-info/auth"
	"github.com/faruqisan/social-info/auth/google"
	"github.com/gin-gonic/gin"
)

type CallbackHandler struct {
	GoogleAPI auth.Auth
}

func NewCallbackHandler() *CallbackHandler {

	gAPI := google.NewGoogleAPI()

	return &CallbackHandler{gAPI}

}

func (cb *CallbackHandler) HandleGoogleCallback(c *gin.Context) {

	code := c.Query("code")
	errAuth := c.Query("error")
	if errAuth != "" {
		return
	}

	cb.GoogleAPI.GetAccessToken(code)

}
