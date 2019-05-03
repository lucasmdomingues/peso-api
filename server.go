package main

import (
	"os"
	"peso-api/src/handlers"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	app := iris.New()

	app.Use(recover.New())
	app.Use(logger.New())

	app.Get("/", handlers.IndexHandler)
	app.Get("/imc", handlers.ImcHandler)
	app.Get("/weights", handlers.WeightsHandler)

	app.Run(iris.Addr(":"+port), iris.WithoutServerError(iris.ErrServerClosed))
}
