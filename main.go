package main

import (
	"mime"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/shimoju/yosasou/converter"
)

func main() {
	converter.Initialize()
	defer converter.Terminate()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "YOSASOU",
		})
	})

	router.GET("/images/:uri", func(c *gin.Context) {
		uri := c.Param("uri")
		contentType := mime.TypeByExtension(path.Ext(uri))
		image := converter.Image{FileName: uri, ContentType: contentType}

		c.Data(http.StatusOK, image.ContentType, image.Resize())
	})

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3301"
	}

	router.Run(":" + port)
}
