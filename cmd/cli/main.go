package main

import (
	"fmt"

	"github.com/sayarb/expense-tracker/internals/config"
	"github.com/sayarb/expense-tracker/internals/creds"
	"github.com/sayarb/expense-tracker/pkg/auth"
	sheetsutil "github.com/sayarb/expense-tracker/pkg/sheetsutils"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

var rootCmd = &cobra.Command{
	Use:   "xpans",
	Short: "Expense Tracker CLI",
	Long:  `Expense Tracker CLI`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Google Login to CLI application",
	Long:  "Google Login to CLI application",
	Run: func(cmd *cobra.Command, args []string) {
		isLoggedIn := auth.IsUserLoggedIn()
		authConfig, err := config.GetAuthConfig()
		if !isLoggedIn {
			if err != nil {
				panic(err)
			}
			fmt.Println("Head over to http://localhost:8000/auth to login")
			auth.LoginUser(authConfig)
			fmt.Println("Login successful")
		} else {
			fmt.Println("User is already logged in")
		}
	},
}

var createSpreadsheetCmd = &cobra.Command{
	Use:   "init [name]",
	Short: "Create a new spreadsheet",
	Long:  "Create a new spreadsheet",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("requires 1 argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		authConfig, err := config.GetAuthConfig()
		if err != nil {
			panic(err)
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
		spConf := &sheetsutil.SpreadsheetConfig{Name: args[0]}
		sheetService, err := sheetsutil.CreateSpreadsheet(spConf, authConfig, token)

		if err != nil {
			panic(err)
		}

		spreadsheetId := sheetService.Spreadsheet.SpreadsheetId
		creds.Set(creds.KeyringSpreadsheet, spreadsheetId)
		fmt.Println("Spreadsheet created successfully")
	},
}

func init() {
	config.LoadEnv()
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(createSpreadsheetCmd)
}

func main() {
	rootCmd.Execute()
}
