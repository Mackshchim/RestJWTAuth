package auth_endpoint

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthService interface {
	Authorize(authData map[string]string) (access, refresh string, err error)
}

type JWTAuthEndpoint struct {
	service AuthService
}

func NewJWTAuthEndpoint(service AuthService) *JWTAuthEndpoint {
	return &JWTAuthEndpoint{service: service}
}

func (e *JWTAuthEndpoint) AuthHandler(c echo.Context) error {
	data := map[string]string{
		"username": c.FormValue("username"),
		"password": c.FormValue("password"),
	}
	access, refresh, err := e.service.Authorize(data)
	if err != nil {
		respBody, _ := json.Marshal(err)
		return c.JSON(http.StatusBadRequest, respBody)
	}

	respData := map[string]string{
		"access":  access,
		"refresh": refresh,
	}

	respBody, err := json.Marshal(respData)
	if err != nil {
		respBody, err = json.Marshal(err)
		return c.JSON(http.StatusBadRequest, respBody)
	}

	return c.JSON(http.StatusAccepted, respBody)
}
