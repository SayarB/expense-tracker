package creds

import (
	"github.com/zalando/go-keyring"
)

const (
	KeyringServiceName  = "expense-tracker"
	KeyringAccessToken  = "access_token"
	KeyringRefreshToken = "refresh_token"
	KeyringSpreadsheet  = "spreadsheet"
	KeyringToken        = "token"
)

func Get(k string) (string, error) {
	// fmt.Printf("Getting \nkey: %s\n", k)
	v, err := keyring.Get(KeyringServiceName, k)
	if err != nil {
		return "", nil
	}
	return v, nil
}

func Set(k string, v string) error {
	// fmt.Printf("Saving \nkey: %s\nvalue: %s\n", k, v)
	return keyring.Set(KeyringServiceName, k, v)
}

func Delete(k string) error {
	return keyring.Delete(KeyringServiceName, k)
}

func DeleteAll() error {
	return keyring.DeleteAll(KeyringServiceName)
}
