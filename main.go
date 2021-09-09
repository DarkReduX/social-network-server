package main

import (
	"database/sql"
	"fmt"
	"github.com/DarkReduX/social-network-server/internal/config"
	"github.com/DarkReduX/social-network-server/internal/handler"
	"github.com/DarkReduX/social-network-server/internal/repository"
	"github.com/DarkReduX/social-network-server/internal/service"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

const (
	profilePath = "/profile"
)

func main() {
	log.SetLevel(log.DebugLevel)

	//init configuration from env variables
	postgresCfg := config.NewPostgresConfig()

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

	//init repositories
	profileRepository := repository.NewProfileRepository(db)

	//init services
	profileService := service.NewProfileService(profileRepository)

	//init handlers
	profileHandler := handler.NewProfileHandler(profileService)

	e := echo.New()

	//profile routes
	e.GET(profilePath, profileHandler.Get)
	e.POST(profilePath, profileHandler.Create)
	e.PUT(profilePath, profileHandler.Update)
	e.DELETE(profilePath, profileHandler.Delete)

	e.Logger.Fatal(e.Start(":1323"))
}
