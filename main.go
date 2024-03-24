package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/erenyusufduran/student-lesson/data"
	"github.com/erenyusufduran/student-lesson/db"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

const webPort = ":80"

type Config struct {
	DB     *gorm.DB
	Models data.Models
}

func main() {
	gormDB, err := db.InitDb()
	if err != nil {
		log.Panic("Can not connect to mySQL!")
	}

	app := Config{
		DB:     gormDB,
		Models: data.New(gormDB),
	}

	s := http.Server{
		Addr:    webPort,
		Handler: app.routes(),
	}

	fmt.Printf("Server is started on %s", webPort)

	if err := s.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}

func (app *Config) routes() *echo.Echo {
	e := echo.New()

	e.POST("/create", app.registerStudent)
	e.POST("/login", app.login)

	e.POST("/plans", app.createPlan, app.authMiddleware())

	return e
}
