package auth

import (
	"net/http"

	"github.com/sayarb/expense-tracker/internals/config"
)

func (authServer *AuthServer) RegisterAuthRoutes(mux *http.ServeMux, configParams *config.AuthConfig) {

	mux.HandleFunc("/auth", AuthHandler(configParams))
	mux.HandleFunc("/auth/callback", CallbackHandler(configParams))
	mux.HandleFunc("/auth/success", SuccessHandler(func() {
		authServer.Close()
	}))
}
