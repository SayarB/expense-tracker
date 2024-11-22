package creds

import "github.com/zalando/go-keyring"

const (
	keyringServiceName  = "expense-tracker"
	keyringAccessToken  = "access_token"
	keyringRefreshToken = "refresh_token"
)

func GetAccessToken() (string, error) {
	token, err := keyring.Get(keyringServiceName, keyringAccessToken)
	if err != nil {
		return "", err
	}
	return token, nil
}

func SetAccessToken(token string) error {
	err := keyring.Set(keyringServiceName, keyringAccessToken, token)
	if err != nil {
		return err
	}
	return nil
}

func GetRefreshToken() (string, error) {
	token, err := keyring.Get(keyringServiceName, keyringRefreshToken)
	if err != nil {
		return "", err
	}
	return token, nil
}

func SetRefreshToken(token string) error {
	err := keyring.Set(keyringServiceName, keyringRefreshToken, token)
	if err != nil {
		return err
	}
	return nil
}

func DeleteAll() error {
	err := keyring.Delete(keyringServiceName, keyringAccessToken)
	if err != nil {
		return err
	}
	err = keyring.Delete(keyringServiceName, keyringRefreshToken)
	if err != nil {
		return err
	}
	return nil
}
