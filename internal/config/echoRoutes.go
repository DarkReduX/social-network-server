package config

import (
	"github.com/DarkReduX/social-network-server/internal/handler"
	localMiddleware "github.com/DarkReduX/social-network-server/internal/middleware"
	"github.com/DarkReduX/social-network-server/internal/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

const (
	authPath    = "/auth"
	profilePath = "/profile"
	friendPath  = "/friend"
)

func NewEchoWithRoutes(e *echo.Echo, jwtConfig middleware.JWTConfig, profileHandler *handler.ProfileHandler, authHandler *handler.AuthHandler, friendHandler *handler.FriendHandler, tokenRepository *repository.AuthRepository) *echo.Echo {
	//swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	//auth routes
	e.POST(authPath, authHandler.Login)
	e.DELETE(authPath, authHandler.Logout, localMiddleware.AuthenticateToken(jwtConfig, tokenRepository))

	//profile routes
	e.GET(profilePath, profileHandler.Get)
	e.POST(profilePath, profileHandler.Create)
	e.PUT(profilePath, profileHandler.Update, localMiddleware.AuthenticateToken(jwtConfig, tokenRepository))
	e.DELETE(profilePath, profileHandler.Delete, localMiddleware.AuthenticateToken(jwtConfig, tokenRepository))

	//friend routes
	e.POST(friendPath, friendHandler.AddFriendRequest, localMiddleware.AuthenticateToken(jwtConfig, tokenRepository))
	e.PUT(friendPath, friendHandler.ProcessFriendRequest, localMiddleware.AuthenticateToken(jwtConfig, tokenRepository))

	return e
}
