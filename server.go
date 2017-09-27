package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/shimoju/yosasou/converter"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "YOSASOU",
		})
	})

	router.GET("/images/:uri", func(c *gin.Context) {
		uri := c.Param("uri")
		image := converter.Image{FileName: uri, ContentType: "image/png"}

		c.Data(http.StatusOK, image.ContentType, image.Resize())
	})

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3301"
	}

	router.Run(":" + port)
}
