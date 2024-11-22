package main

import (
	"fmt"

	"github.com/sayarb/expense-tracker/internals/config"
	"github.com/sayarb/expense-tracker/internals/creds"
	"github.com/sayarb/expense-tracker/pkg/auth"
	sheetsutil "github.com/sayarb/expense-tracker/pkg/sheetsutils"
	"golang.org/x/oauth2"
)

func init() {
	config.LoadEnv()

}

func main() {
	loggedIn := auth.IsUserLoggedIn()
	authConfig, err := config.GetAuthConfig()
	if err != nil {
		panic(err)
	}

	if !loggedIn {
		auth.LoginUser(authConfig)
	} else {
		fmt.Println("User is already logged in")
	}

	accessToken, err := creds.Get(creds.KeyringAccessToken)
	if err != nil {
		panic(err)
	}

	refreshToken, err := creds.Get(creds.KeyringRefreshToken)
	if err != nil {
		panic(err)
	}

	token := &oauth2.Token{AccessToken: accessToken, RefreshToken: refreshToken, TokenType: "Bearer"}

	sheet, err := sheetsutil.CreateSpreadsheet(&sheetsutil.SpreadsheetConfig{Name: "Expenses234234"}, authConfig, token)
	if err != nil {
		panic(err)
	}

	creds.Set(creds.KeyringSpreadsheet, sheet.SpreadsheetId)
}
