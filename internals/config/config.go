package config

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type AuthURL struct {
	URL          string
	State        string
	CodeVerifier string
}

func (u *AuthURL) String() string {
	return u.URL
}

type AuthConfig struct {
	CodeChallenge string
	OAuthConfig   *oauth2.Config
	AuthUrl       *AuthURL
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func GetAuthConfig() (*AuthConfig, error) {

	state, stateErr := randomBytesInHex(24)
	if stateErr != nil {
		return nil, fmt.Errorf("Could not generate random state: %v", stateErr)
	}

	config := &oauth2.Config{
		ClientID:    os.Getenv("GCP_CLIENT_ID"),
		RedirectURL: os.Getenv("CALLBACK_URL"),
		Scopes: []string{
			"https://www.googleapis.com/auth/spreadsheets",
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}

	codeVerifier := oauth2.GenerateVerifier()

	authUrl := config.AuthCodeURL(
		state,
		oauth2.S256ChallengeOption(codeVerifier),
		oauth2.AccessTypeOffline,
		oauth2.ApprovalForce,
	)

	return &AuthConfig{OAuthConfig: config, AuthUrl: &AuthURL{URL: authUrl, State: state, CodeVerifier: codeVerifier}}, nil
}

func randomBytesInHex(count int) (string, error) {
	buf := make([]byte, count)
	_, err := io.ReadFull(rand.Reader, buf)
	if err != nil {
		return "", fmt.Errorf("Could not generate %d random bytes: %v", count, err)
	}

	return hex.EncodeToString(buf), nil
}
