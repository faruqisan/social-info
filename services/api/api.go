package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/faruqisan/social-info/auth/google"
	"github.com/faruqisan/social-info/services/api/middlewares"
	"github.com/faruqisan/social-info/src/youtube"
)

func RegisterAPI() {

	g := google.NewGoogleAPI()

	http.HandleFunc("/youtube", middlewares.CheckGoogleAccessToken(g, func(w http.ResponseWriter, r *http.Request) {

		if !g.CheckToken() {
			log.Println("Checking token")
			http.Redirect(w, r, g.GetAuthorizeURL(), http.StatusFound)
			return
		}

		gC := g.GetAPIClient()
		youtubeService := youtube.NewYoutubeClient(gC)

		channelInfos := youtubeService.GetChannelInfos("snippet,contentDetails,statistics")

		p, _ := json.Marshal(channelInfos)

		w.Write(p)

	}))

	http.HandleFunc("/callback/google", func(w http.ResponseWriter, r *http.Request) {

		code := r.FormValue("code")
		errAuth := r.FormValue("error")
		if errAuth != "" {
			return
		}
		g.GetAccessToken(code)

		http.Redirect(w, r, "/youtube", http.StatusFound)

	})
}
