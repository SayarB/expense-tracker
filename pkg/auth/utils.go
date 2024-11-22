package auth

import (
	"github.com/sayarb/expense-tracker/internals/config"
	"github.com/sayarb/expense-tracker/internals/creds"
)

func IsUserLoggedIn() bool {
	token, err := creds.Get(creds.KeyringAccessToken)
	if err != nil {
		return false
	}
	return token != ""
}

func LoginUser(config *config.AuthConfig) {
	authServer := NewAuthServer(config)
	authServer.ListenAndServe()
}
