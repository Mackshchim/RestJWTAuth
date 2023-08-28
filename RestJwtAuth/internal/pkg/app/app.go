package app

import (
	"RestJwtAuth/db/client/mongodb"
	"RestJwtAuth/internal/app/endpoint/auth-endpoint"
	refresh_tokens_repository "RestJwtAuth/internal/app/repository/refresh-tokens-repository"
	users_repository "RestJwtAuth/internal/app/repository/users-repository"
	"RestJwtAuth/internal/app/services/jwt-handlers"
	"context"
	"github.com/labstack/echo/v4"
)

const (
	SecretKey = "secret"

	DbName                  = "test_db"
	UsersCollection         = "user"
	RefreshTokensCollection = ""
)

type AuthEndpoint interface {
	AuthHandler(c echo.Context) error
}

type RefreshEndpoint interface {
	RefreshHandler(c echo.Context) error
}

type App struct {
	authEndpoint    AuthEndpoint
	refreshEndpoint RefreshEndpoint
	echoServer      *echo.Echo
}

func New() (*App, error) {
	app := &App{}

	ctx := context.Background()
	client, err := mongodb.NewClient(ctx, "localhost", "27017")
	if err != nil {
		return nil, err
	}

	usersRepo, err := users_repository.NewCRUDRepository(client)
	rtRepo := refresh_tokens_repository.New(client)
	app.authEndpoint = auth_endpoint.NewJWTAuthEndpoint(jwt_handlers.NewJWTAuthService(usersRepo, rtRepo, SecretKey))

	app.refreshEndpoint = auth_endpoint.NewRefresh(jwt_handlers.NewRefreshService(rtRepo))

	app.echoServer = echo.New()

	app.echoServer.POST("/auth", app.authEndpoint.AuthHandler)
	app.echoServer.POST("/auth/refresh", app.refreshEndpoint.RefreshHandler)

	return app, nil
}

func (a *App) Run(address string) error {
	err := a.echoServer.Start(address)
	return err
}
