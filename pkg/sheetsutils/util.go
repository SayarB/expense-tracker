package sheetsutil

import (
	"context"
	"fmt"
	"log"

	"github.com/sayarb/expense-tracker/internals/config"
	"golang.org/x/oauth2"
	"google.golang.org/api/option"
	sheets "google.golang.org/api/sheets/v4"
)

type SpreadsheetConfig struct {
	Name string
}

type SpreadsheetService struct {
	Service     *sheets.Service
	Spreadsheet *sheets.Spreadsheet
}

func createSheetsService(config *config.AuthConfig, token *oauth2.Token) (*sheets.Service, error) {
	client := config.OAuthConfig.Client(context.Background(), token)
	SheetsService, err := sheets.NewService(context.Background(), option.WithHTTPClient(client))
	return SheetsService, err
}

func CreateSpreadsheet(spConf *SpreadsheetConfig, authConfig *config.AuthConfig, token *oauth2.Token) (*SpreadsheetService, error) {
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
		return nil, err
	}

	_, err = service.Spreadsheets.Values.Append(newSpreadsheet.SpreadsheetId, "A1:B1", &sheets.ValueRange{MajorDimension: "ROWS", Values: [][]interface{}{{"A", "B"}}}).ValueInputOption("USER_ENTERED").Do()

	if err != nil {
		log.Println(err.Error())
	}

	fmt.Printf("Spreadsheet created with ID: %s\n", newSpreadsheet.SpreadsheetId)
	return &SpreadsheetService{Service: service, Spreadsheet: newSpreadsheet}, nil
}
