package main

import (
	"fmt"

	"github.com/sayarb/expense-tracker/internals/config"
	"github.com/sayarb/expense-tracker/pkg/auth"
	"github.com/spf13/cobra"
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
		if !isLoggedIn {
			authConfig, err := config.GetAuthConfig()
			if err != nil {
				panic(err)
			}
			fmt.Println("Head over to http://localhost:8000/auth to login")
			auth.LoginUser(authConfig)
			fmt.Println("Login successful")
			return
		} else {
			fmt.Println("User is already logged in")
		}
	},
}

func init() {
	config.LoadEnv()
	rootCmd.AddCommand(loginCmd)
}

func main() {
	rootCmd.Execute()
}
