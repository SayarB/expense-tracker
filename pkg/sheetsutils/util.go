package sheetsutil

import (
	"context"
	"fmt"
	"log"

	"github.com/sayarb/expense-tracker/internals/config"
	"github.com/sayarb/expense-tracker/internals/storage"
	"golang.org/x/oauth2"
	"google.golang.org/api/option"
	sheets "google.golang.org/api/sheets/v4"
)

var (
	headerRow = [][]interface{}{{"Date", "Amount", "Reciever", "Reason"}}
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

	_, err = service.Spreadsheets.Values.Append(newSpreadsheet.SpreadsheetId, "A1:B1", &sheets.ValueRange{MajorDimension: "ROWS", Values: headerRow}).ValueInputOption("USER_ENTERED").Do()

	if err != nil {
		log.Println(err.Error())
	}

	fmt.Printf("Spreadsheet created with ID: %s\n", newSpreadsheet.SpreadsheetId)
	storage.SetNumberOfRecords(0)
	return &SpreadsheetService{Service: service, Spreadsheet: newSpreadsheet}, nil
}

func GetSpreadsheet(spreadsheetId string, authConfig *config.AuthConfig, token *oauth2.Token) (*SpreadsheetService, error) {
	service, err := createSheetsService(authConfig, token)
	if err != nil {
		return nil, err
	}
	spreadsheet, err := service.Spreadsheets.Get(spreadsheetId).Do()
	if err != nil {
		return nil, err
	}
	return &SpreadsheetService{Service: service, Spreadsheet: spreadsheet}, nil
}

func (service *SpreadsheetService) GetHeaders() (*sheets.ValueRange, error) {
	rangeName := "A1:D1"
	return service.Service.Spreadsheets.Values.Get(service.Spreadsheet.SpreadsheetId, rangeName).Do()
}

func (service *SpreadsheetService) AddExpense(row []interface{}) error {
	rowNumber := storage.GetNumberOfRecords()
	rangeName := fmt.Sprintf("A%d:D%d", rowNumber+1, rowNumber+1)
	_, err := service.Service.Spreadsheets.Values.Append(service.Spreadsheet.SpreadsheetId, rangeName, &sheets.ValueRange{MajorDimension: "ROWS", Values: [][]interface{}{row}}).ValueInputOption("USER_ENTERED").Do()
	if err != nil {
		return err
	}
	storage.SetNumberOfRecords(storage.GetNumberOfRecords() + 1)
	return nil
}
