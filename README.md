# Expense Tracker

Expense Tracker is a CLI application that allows users to track their expenses. It uses Google Sheets as the backend and provides a simple command-line interface to interact with the application.

## Features

- Login with Google OAuth
- Create a new spreadsheet
- Add expenses to the spreadsheet
- Retrieve expenses from the spreadsheet

## Getting Started

To get started, you need to create a Google Cloud Platform project and enable the Google Sheets API. Once you have done that, you can follow these steps to set up the application:

1. Clone the repository:

```bash
git clone https://github.com/sayarb/expense-tracker.git
```

2. Install the dependencies:

```bash
go get .
```

3. Create a new Google Cloud Platform project and enable the Google Sheets API.

4. Create a new OAuth client ID for the application. Follow the instructions in the [Google OAuth documentation](https://developers.google.com/identity/protocols/oauth2) to create a new client ID. (Remember to use UWP type of application. This removes the need for a Client Secret During the PKCE flow)

5. Create a new `.env` file in the root directory of the project and add the following environment variables:

```bash
PORT=8000
GCP_CLIENT_SECRET=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
GCP_CLIENT_ID=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
CALLBACK_URL=http://localhost:8000/auth/callback
```

6. Run the application:

```bash
go run .
```

7. Build the executable:

```bash
go build ./cmd/cli -o xpans.exe
```

8. Run the executable:

```bash
./xpans
```

## Commands

### Login

The `login` command allows you to log in to the application using Google OAuth. It will open a browser window and prompt you to log in to your Google account. After you have logged in, the application will store your tokens in the `keyring`.

```bash
./xpans login
```

### Create Spreadsheet

The `init` command allows you to create a new spreadsheet in Google Sheets. It will prompt you to enter a name for the spreadsheet.

```bash
./xpans init "Expenses"
```

### Add Expense

The `add` command allows you to add an expense to the spreadsheet. It will prompt you to enter the date, amount, recipient, and reason for the expense.

```bash
./xpans add 1000 "John Doe" "Lunch with friends"
```
