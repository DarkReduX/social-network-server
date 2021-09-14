package main

import (
	"database/sql"
	"fmt"
	"github.com/DarkReduX/social-network-server/internal/config"
	"github.com/DarkReduX/social-network-server/internal/handler"
	"github.com/DarkReduX/social-network-server/internal/models"
	"github.com/DarkReduX/social-network-server/internal/repository"
	"github.com/DarkReduX/social-network-server/internal/service"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

const (
	authPath    = "/auth"
	profilePath = "/profile"
	friendPath  = "/friend"
)

func main() {
	log.SetLevel(log.DebugLevel)

	//init configuration from env variables
	postgresCfg := config.NewPostgresConfig()
	redisCfg := config.NewRedisConfig()

	//init databases
	db, err := sql.Open("postgres", fmt.Sprintf("port=%s host=%s user=%s password=%s dbname=%s sslmode=%s",
		postgresCfg.Port,
		postgresCfg.Host,
		postgresCfg.User,
		postgresCfg.Password,
		postgresCfg.DbName,
		postgresCfg.SslMode,
	))
	if err != nil {
		log.WithFields(log.Fields{
			"file": "main.go",
			"func": "main",
		}).Errorf("Couldn't open db connection: %v", err)
	}
	if err = db.Ping(); err != nil {
		log.Errorf("%v", err)
	}
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Username: redisCfg.Username,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})
	//init repositories
	profileRepository := repository.NewProfileRepository(db)
	authRepository := repository.NewAuthRepository(redisClient)
	friendRepository := repository.NewFriendRepository(db)

	//init services
	profileService := service.NewProfileService(profileRepository)
	authService := service.NewAuthService(profileRepository, authRepository)
	friendService := service.NewFriendService(friendRepository)

	//init handlers
	profileHandler := handler.NewProfileHandler(profileService)
	authHandler := handler.NewAuthHandler(authService)
	friendHandler := handler.NewFriendHandler(friendService)

	jwtConfig := middleware.JWTConfig{
		SigningKey: []byte("user"),
		Claims:     &models.CustomClaims{},
	}

	e := echo.New()

	//auth routes
	e.POST(authPath, authHandler.Login)
	e.DELETE(authPath, authHandler.Logout, authService.AuthenticateTokenMiddleware(jwtConfig))

	//profile routes
	e.GET(profilePath, profileHandler.Get)
	e.POST(profilePath, profileHandler.Create)
	e.PUT(profilePath, profileHandler.Update, authService.AuthenticateTokenMiddleware(jwtConfig))
	e.DELETE(profilePath, profileHandler.Delete, authService.AuthenticateTokenMiddleware(jwtConfig))

	//friend routes
	e.POST(friendPath, friendHandler.AddFriendRequest, authService.AuthenticateTokenMiddleware(jwtConfig))
	e.DELETE(friendPath, friendHandler.DeleteFriend, authService.AuthenticateTokenMiddleware(jwtConfig))
	e.PUT(friendPath, friendHandler.SubmitFriendRequest, authService.AuthenticateTokenMiddleware(jwtConfig))

	e.Logger.Fatal(e.Start(":1323"))
}
