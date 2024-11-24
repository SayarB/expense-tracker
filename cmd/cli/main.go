package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/sayarb/expense-tracker/internals/config"
	"github.com/sayarb/expense-tracker/internals/creds"
	"github.com/sayarb/expense-tracker/internals/storage"
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
		tokenJson, err := creds.Get(creds.KeyringToken)
		if err != nil {
			panic(err)
		}
		token := &oauth2.Token{}
		json.Unmarshal([]byte(tokenJson), token)

		spId, err := creds.Get(creds.KeyringSpreadsheet)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Found Spreadsheet ID: %s\n", spId)

		sheetService, err := sheetsutil.GetSpreadsheet(spId, authConfig, token)
		if err == nil && sheetService.Spreadsheet.Properties.Title == args[0] {
			fmt.Println("Spreadsheet already exists")
			value, err := sheetService.GetHeaders()
			if err != nil {
				panic(err)
			}
			fmt.Print(value.Values)
			return
		}

		spConf := &sheetsutil.SpreadsheetConfig{Name: args[0]}
		sheetService, err = sheetsutil.CreateSpreadsheet(spConf, authConfig, token)

		if err != nil {
			panic(err)
		}

		spreadsheetId := sheetService.Spreadsheet.SpreadsheetId
		creds.Set(creds.KeyringSpreadsheet, spreadsheetId)
		fmt.Println("Spreadsheet created successfully")
	},
}

var addExpenseCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an expense to the spreadsheet",
	Long:  "Add an expense to the spreadsheet",
	Run: func(cmd *cobra.Command, args []string) {
		authConfig, err := config.GetAuthConfig()
		if err != nil {
			panic(err)
		}
		tokenJson, err := creds.Get(creds.KeyringToken)
		if err != nil {
			panic(err)
		}
		token := &oauth2.Token{}
		json.Unmarshal([]byte(tokenJson), token)

		spId, err := creds.Get(creds.KeyringSpreadsheet)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Found Spreadsheet ID: %s\n", spId)

		sheetService, err := sheetsutil.GetSpreadsheet(spId, authConfig, token)
		if err != nil {
			panic(err)
		}

		var row []interface{}
		row = append(row, time.Now().Format(time.DateTime))
		row = append(row, args[0])
		row = append(row, args[1])
		row = append(row, args[2])

		err = sheetService.AddExpense(row)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Expense added successfully\nCheckout the changes at %s\n", sheetService.Spreadsheet.SpreadsheetUrl)
	},
}

func init() {
	config.LoadEnv()
	storage.ReadConfigFile()
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(createSpreadsheetCmd)
	rootCmd.AddCommand(addExpenseCmd)
}

func main() {
	rootCmd.Execute()
}
