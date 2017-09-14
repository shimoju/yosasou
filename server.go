package main

import (
	"net/http"
	"os"

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

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3301"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
