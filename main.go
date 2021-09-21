package main

import (
	"database/sql"
	"fmt"
	_ "github.com/DarkReduX/social-network-server/docs"
	"github.com/DarkReduX/social-network-server/internal/config"
	"github.com/DarkReduX/social-network-server/internal/handler"
	"github.com/DarkReduX/social-network-server/internal/models"
	"github.com/DarkReduX/social-network-server/internal/repository"
	"github.com/DarkReduX/social-network-server/internal/service"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

// @title Social Network Server
// @version 1.0
// @description HTTP server for social network.

// @host localhost:1323
// @BasePath /

// @securityDefinitions.apikey BearerToken
// @in header
// @name Authorization
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

	e = config.NewEchoWithRoutes(e, jwtConfig, profileHandler, authHandler, friendHandler, authRepository)

	// enable jaeger middleware
	c := jaegertracing.New(e, nil)
	defer func() {
		if jErr := c.Close(); jErr != nil {
			log.Errorf("Error while closing tracer: %v", jErr)
		}
	}()
	e.Logger.Fatal(e.Start(":1323"))
}
