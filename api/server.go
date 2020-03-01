package api

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Serve launches the web client
// Using Gin & Vue.js
func Serve(port int) {
	httpPort := ":" + strconv.Itoa(port)

	router := gin.Default()

	router.Group("/api").
		GET("/", healthHandler).
		GET("/numbers", getAllNumbers).
		GET("/numbers/:number/scan/local", ValidateScanURL, localScan).
		GET("/numbers/:number/scan/numverify", ValidateScanURL, numverifyScan).
		GET("/numbers/:number/scan/googlesearch", ValidateScanURL, googleSearchScan)

	dir, _ := os.Getwd()
	assetsPath := dir + "/client/dist"

	router.Group("/").
		Static("/js", assetsPath+"/js").
		Static("/css", assetsPath+"/css").
		Static("/img", assetsPath+"/img")

	router.StaticFile("/favicon.ico", assetsPath+"/favicon.ico")
	router.LoadHTMLFiles(assetsPath + "/index.html")

	router.GET("/", func(c *gin.Context) {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.Use(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"success": false,
			"message": "Resource not found",
		})
	})

	router.Run(httpPort)
}