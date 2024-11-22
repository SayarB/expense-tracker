package sheetsutil

import (
	"context"
	"fmt"

	"github.com/sayarb/expense-tracker/internals/config"
	"golang.org/x/oauth2"
	"google.golang.org/api/option"
	sheets "google.golang.org/api/sheets/v4"
)

type SpreadsheetConfig struct {
	Name string
}

func createSheetsService(config *config.AuthConfig, token *oauth2.Token) (*sheets.Service, error) {

	client := config.OAuthConfig.Client(context.Background(), token)

	return sheets.NewService(context.Background(), option.WithHTTPClient(client))
}

func CreateSpreadsheet(spConf *SpreadsheetConfig, authConfig *config.AuthConfig, token *oauth2.Token) error {
	spreadsheet := &sheets.Spreadsheet{
		Properties: &sheets.SpreadsheetProperties{
			Title: spConf.Name,
		},
		Sheets: []*sheets.Sheet{{
			Properties: &sheets.SheetProperties{
				Title: "Sheet1",
			},
		}},
	}
	service, err := createSheetsService(authConfig, token)

	newSpreadsheet, err := service.Spreadsheets.Create(spreadsheet).Do()

	if err != nil {
		return err
	}
	
	fmt.Println(newSpreadsheet.SpreadsheetId)
	return nil
}
