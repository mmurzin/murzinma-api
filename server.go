package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

const PORT = 8080

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "{\"gretting\": \"Hello, World!!\"}")
	})
	e.Logger.Fatal(e.Start(":" + strconv.FormatInt(PORT, 10)))
}
