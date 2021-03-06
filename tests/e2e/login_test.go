package e2e

import (
	"net/http"
	"testing"

	"github.com/MTES-MCT/filharmonic-api/authentication"
	"github.com/MTES-MCT/filharmonic-api/httpserver"
	"github.com/MTES-MCT/filharmonic-api/tests"
)

func TestLoginSuccessful(t *testing.T) {
	e, close, sso := tests.InitWithSso(t)
	defer close()

	sso.On("ValidateTicket", "ticket-exploitant1").Return("exploitant1@filharmonic.com", nil)

	loginResult := e.POST("/login").WithJSON(&httpserver.LoginHTTPRequest{Ticket: "ticket-exploitant1"}).
		Expect().Status(http.StatusOK).JSON().Object()
	loginResult.ContainsKey("token")
}

func TestLoginFailed(t *testing.T) {
	e, close, sso := tests.InitWithSso(t)
	defer close()

	sso.On("ValidateTicket", "invalid-ticket").Return("", authentication.ErrUnauthorized)

	e.POST("/login").WithJSON(&httpserver.LoginHTTPRequest{Ticket: "invalid-ticket"}).
		Expect().Status(http.StatusUnauthorized)
}

func TestAuthenticateSuccessful(t *testing.T) {
	e, close := tests.Init(t)
	defer close()

	e.POST("/authenticate").WithJSON(&httpserver.AuthenticateHTTPRequest{Token: tests.UserSessions[1]}).
		Expect().Status(http.StatusOK)
}

func TestAuthenticateFailed(t *testing.T) {
	e, close := tests.Init(t)
	defer close()

	e.POST("/authenticate").WithJSON(&httpserver.AuthenticateHTTPRequest{Token: "invalid-token"}).
		Expect().Status(http.StatusUnauthorized)
}

func TestLogout(t *testing.T) {
	e, close := tests.Init(t)
	defer close()

	tests.AuthExploitant(e.POST("/logout")).
		Expect().Status(http.StatusOK)

	e.POST("/authenticate").WithJSON(&httpserver.AuthenticateHTTPRequest{Token: tests.UserSessions[1]}).
		Expect().Status(http.StatusUnauthorized)
}
