package auth

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sayarb/expense-tracker/internals/config"
	"github.com/sayarb/expense-tracker/internals/creds"
	"golang.org/x/oauth2"
)

type UserDataSchema struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Picture       string `json:"picture"`
}
type AuthURL struct {
	URL          string
	State        string
	CodeVerifier string
}

func (u *AuthURL) String() string {
	return u.URL
}

var oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func AuthHandler(conf *config.AuthConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "GET" && r.Method != "" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var expiration = time.Now().Add(365 * 24 * time.Hour)

		cookie := &http.Cookie{Name: "oauthstate", Value: conf.AuthUrl.State, Expires: expiration}

		authUrl := conf.AuthUrl

		http.SetCookie(w, cookie)
		http.Redirect(w, r, authUrl.String(), http.StatusTemporaryRedirect)
	}
}

func CallbackHandler(config *config.AuthConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "GET" && r.Method != "" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		oauthState, err := r.Cookie("oauthstate")

		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/", fiber.StatusPermanentRedirect)
		}

		if r.FormValue("state") != oauthState.Value {
			log.Println("invalid oauth google state")
			http.Redirect(w, r, "/", fiber.StatusPermanentRedirect)
		}

		code := r.FormValue("code")

		token, err := config.OAuthConfig.Exchange(context.Background(), code, oauth2.VerifierOption(config.AuthUrl.CodeVerifier))
		if err != nil {
			log.Println(err.Error())
		}

		// response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
		// if err != nil {
		// 	log.Println(err.Error())
		// }
		// defer response.Body.Close()
		// contents, err := io.ReadAll(response.Body)
		// if err != nil {
		// 	log.Println(err.Error())
		// }

		// userData := &UserDataSchema{}
		// json.Unmarshal([]byte(contents), userData)

		tokenJson, err := json.Marshal(token)
		if err != nil {
			log.Println(err.Error())
		}
		creds.Set(creds.KeyringToken, string(tokenJson))
		http.Redirect(w, r, "/auth/success", fiber.StatusPermanentRedirect)
	}
}

func SuccessHandler(callback func()) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		callback()
	}
}
