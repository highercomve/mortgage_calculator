package server

import (
	"net/http"

	"github.com/highercomve/mortgage_calculator/api"
	"github.com/highercomve/mortgage_calculator/docs"
	"github.com/highercomve/mortgage_calculator/modules/validation"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"

	echoSwagger "github.com/swaggo/echo-swagger"
)

func Start(serverAddress string) {
	docs.SwaggerInfo.BasePath = "/api/v1/"
	docs.SwaggerInfo.Schemes = []string{"https"}
	docs.SwaggerInfo.Host = viper.GetString("host-url")

	e := echo.New()
	e.Validator = validation.GetValidator()

	// Pre request middlewares
	e.Pre(middleware.RemoveTrailingSlash())

	corsConfig := middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodHead,
			http.MethodPut,
			http.MethodPatch,
			http.MethodPost,
			http.MethodDelete,
			http.MethodOptions,
		},
	}

	// Set Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	e.Use(middleware.RequestID())
	e.Use(middleware.CORSWithConfig(corsConfig))

	if viper.GetBool("debug") {
		e.Debug = true
	}

	// Metrics
	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)

	// Load all api services
	api.LoadAPI(e)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(serverAddress))
}
