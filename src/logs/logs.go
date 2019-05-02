package logs

import (
	"log"

	"github.com/kataras/iris"
)

func Info(ctx iris.Context, data interface{}, message string) {
	ctx.JSON(iris.Map{"data": data, "message": message})
}

func Error(ctx iris.Context, err error, message string) {
	log.Println(err.Error())
	ctx.JSON(iris.Map{"error": message})
}
