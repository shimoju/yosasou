package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/shimoju/yosasou/converter"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/:uri", func(c echo.Context) error {
		// uri := c.Param("uri")
		image := converter.ResizeImage()
		return c.Blob(http.StatusOK, "image/png", image)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
