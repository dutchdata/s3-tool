package main

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()
	e.GET("/",func(c echo.Context) (error) {
		return c.String(http.StatusOK, "ok")
	})

	e.GET("/auth", public.accessKeyHandler)
	e.GET("/go", public.recordHandler)
	e.GET("/get", public.downloadHandler)
	e.GET("/check-for-trails",public.trailCheckHandler)
	e.GET("/get-trail-events", public.trailEventHandler)

	e.HideBanner = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Logger.Fatal(e.Start(":8080"))
}
