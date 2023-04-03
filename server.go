package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/acme/autocert"
)

const PORT = 8080

func main() {
	e := echo.New()
	// e.AutoTLSManager.HostPolicy = autocert.HostWhitelist("<DOMAIN>")
	// Cache certificates to avoid issues with rate limits (https://letsencrypt.org/docs/rate-limits)
	e.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")
	//e.Use(middleware.Recover())
	//e.Use(middleware.Logger())
	e.Use(middleware.HTTPSRedirect())
	e.Pre(middleware.HTTPSNonWWWRedirect())
	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, `
			<h1>Welcome to Echo!</h1>
			<h3>TLS certificates automatically installed from Let's Encrypt :)</h3>
		`)
	})

	e.Logger.Fatal(e.StartAutoTLS(":443"))
}
