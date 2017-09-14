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
		image := converter.Image{FileName: "icon.png", ContentType: "image/png"}
		return c.Blob(http.StatusOK, image.ContentType, image.Resize())
	})

	e.Logger.Fatal(e.Start(":1323"))
}
