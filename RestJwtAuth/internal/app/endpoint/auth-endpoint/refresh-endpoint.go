package auth_endpoint

import (
	web_tokens "RestJwtAuth/internal/pkg/utils/web-tokens"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type RefreshService interface {
	Refresh(access, refresh string) (newAccess, newRefresh string, err error)
}

type JWTRefreshEndpoint struct {
	service RefreshService
}

func NewRefresh(service RefreshService) *JWTRefreshEndpoint {
	return &JWTRefreshEndpoint{
		service: service,
	}
}

func (e *JWTRefreshEndpoint) RefreshHandler(c echo.Context) error {
	if !web_tokens.HasRefreshToken(c.Request().Header) {
		return fmt.Errorf("request has no refresh token")
	}
	refresh, _ := strings.CutPrefix(c.Request().Header.Get(web_tokens.RefreshTokenHeader), web_tokens.TokenPrefix)

	if !web_tokens.HasAccessToken(c.Request().Header) {
		return fmt.Errorf("request has no access token")
	}
	access, _ := strings.CutPrefix(c.Request().Header.Get(web_tokens.AccessTokenHeader), web_tokens.TokenPrefix)

	access, refresh, err := e.service.Refresh(access, refresh)
	if err != nil {
		return err
	}

	respData := map[string]string{
		"access":  access,
		"refresh": refresh,
	}

	if err != nil {
		respBody, _ := json.Marshal(err)
		return c.JSON(http.StatusBadRequest, respBody)
	}
	respBody, err := json.Marshal(respData)
	if err != nil {
		respBody, err = json.Marshal(err)
		return c.JSON(http.StatusBadRequest, respBody)
	}

	return c.JSON(http.StatusAccepted, respBody)
}
