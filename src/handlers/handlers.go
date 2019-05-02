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

func ImcHandler(ctx iris.Context) {

	height, err := strconv.ParseFloat(ctx.FormValue("height"), 64)
	if err != nil {
		logs.Error(ctx, err, "Error on parse height")
		return
	}

	weight, err := strconv.ParseFloat(ctx.FormValue("weight"), 64)
	if err != nil {
		logs.Error(ctx, err, "Error on parse height")
		return
	}

	imc := iris.Map{
		"imc": fmt.Sprintf("%.2f", weight/math.Pow(height, 2)),
	}

	logs.Info(ctx, imc, "IMC calculated with success")
}

func WeightsHandler(ctx iris.Context) {

	height, err := strconv.ParseFloat(ctx.FormValue("height"), 64)
	if err != nil {
		logs.Error(ctx, err, "Error on parse height")
		return
	}

	calcMinWeight := math.Pow(height, 2) * MinWeightValue
	calcMaxWeight := math.Pow(height, 2) * MaxWeightValue

	minWeight := fmt.Sprintf("%.2f", calcMinWeight)
	maxWeight := fmt.Sprintf("%.2f", calcMaxWeight)

	weights := iris.Map{
		"min_weight": minWeight,
		"max_weight": maxWeight,
	}

	logs.Info(ctx, weights, "Weights calculated with success")
}
