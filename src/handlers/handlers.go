package handlers

import (
	"fmt"
	"math"
	"peso-api/src/logs"
	"strconv"

	"github.com/kataras/iris"
)

const MinWeightValue = 18.5
const MaxWeightValue = 24.99

func IndexHandler(ctx iris.Context) {
	ctx.JSON(iris.Map{"Title": "Peso API", "Version": "1.1"})
}

func ImcHandler(ctx iris.Context) {

	height, err := strconv.ParseFloat(ctx.FormValue("height"), 64)
	if err != nil {
		logs.Error(ctx, err, "Error on parse height")
		return
	}

	weight, err := strconv.ParseFloat(ctx.FormValue("weight"), 64)
	if err != nil {
		logs.Error(ctx, err, "Error on parse weight")
		return
	}

	imc := weight / math.Pow(height, 2)

	var situation string

	if imc < 17 {
		situation = "Muito abaixo do peso"
	} else if imc >= 17 || imc <= 18.49 {
		situation = "Abaixo do peso"
	} else if imc >= 18.5 || imc <= 24.99 {
		situation = "Peso normal"
	} else if imc >= 25 || imc <= 29.99 {
		situation = "Acima do peso"
	} else if imc >= 30 || imc <= 34.99 {
		situation = "Obesidade I"
	} else if imc >= 35 || imc <= 39.99 {
		situation = "Obesidade II (severa)"
	} else if imc >= 35 || imc <= 39.99 {
		situation = "Obesidade III (mórbida)"
	}

	data := iris.Map{
		"imc":       fmt.Sprintf("%.2f", imc),
		"situation": situation,
	}

	logs.Info(ctx, data, "IMC calculated with success")
}

func WeightsHandler(ctx iris.Context) {

	height, err := strconv.ParseFloat(ctx.FormValue("height"), 64)
	if err != nil {
		logs.Error(ctx, err, "Error on parse height")
		return
	}

	calcMinWeight := math.Pow(height, 2) * MinWeightValue
	calcMaxWeight := math.Pow(height, 2) * MaxWeightValue

	minWeight := fmt.Sprintf("%.0f", calcMinWeight)
	maxWeight := fmt.Sprintf("%.0f", calcMaxWeight)

	weights := iris.Map{
		"min_weight": minWeight,
		"max_weight": maxWeight,
	}

	logs.Info(ctx, weights, "Weights calculated with success")
}
