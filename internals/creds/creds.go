package creds

import "github.com/zalando/go-keyring"

const (
	KeyringServiceName  = "expense-tracker"
	KeyringAccessToken  = "access_token"
	KeyringRefreshToken = "refresh_token"
	KeyringSpreadsheet  = "spreadsheet"
)

func Get(k string) (string, error) {
	v, err := keyring.Get(KeyringServiceName, k)
	if err != nil {
		return "", nil
	}
	return v, nil
}

func Set(k string, v string) error {
	return keyring.Set(KeyringServiceName, k, v)
}

func Delete(k string) error {
	return keyring.Delete(KeyringServiceName, k)
}

func DeleteAll() error {
	return keyring.DeleteAll(KeyringServiceName)
}
