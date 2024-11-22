package auth

import (
	"fmt"
	"net/http"
	"os"

	"github.com/sayarb/expense-tracker/internals/config"
)

type AuthServer struct {
	Server *http.Server
}

func NewAuthServer(config *config.AuthConfig) *AuthServer {
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))

	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    port,
		Handler: mux,
	}

	authServer := &AuthServer{Server: server}
	authServer.RegisterAuthRoutes(mux, config)

	return authServer
}

func (s *AuthServer) ListenAndServe() error {
	return s.Server.ListenAndServe()
}

func (s *AuthServer) Close() error {
	return s.Server.Close()
}
