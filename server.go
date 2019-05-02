package main

import (
	"peso-api/src/handlers"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {

	app := iris.New()

	app.Use(recover.New())
	app.Use(logger.New())

	app.Get("/", handlers.IndexHandler)
	app.Get("/imc", handlers.ImcHandler)
	app.Get("/weights", handlers.WeightsHandler)

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
