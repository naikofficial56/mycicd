package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/sum", getSum)

	e.Logger.Fatal(e.Start(":1323"))

}

func getSum(ctx echo.Context) error {

	// number1 and number 2 are compulsory for this functions
	num1Str := ctx.QueryParam("num1")

	if num1Str == "" {
		return ctx.String(http.StatusOK, "num1 is missing")
	}

	num1, err := strconv.Atoi(num1Str)
	if err != nil {
		return ctx.String(http.StatusOK, "Invalid format of num1")
	}

	result := add(num1)
	return ctx.JSON(http.StatusOK, result)
}
