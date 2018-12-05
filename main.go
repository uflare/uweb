package main

import (
	"log"
	"os"

	"github.com/go-playground/validator"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.HideBanner = true
	e.Validator = &Validator{validator: validator.New()}

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{Level: 9}))
	e.Use(middleware.Recover())

	globalContext.Store("echo", e)

	if err := BootPlugins(); err != nil {
		log.Fatal(err.Error())
	}

	log.Fatal(e.Start(os.Getenv("SERVER_LISTEN_ADDR")))
}
