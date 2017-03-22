package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/ChimeraCoder/anaconda"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello from the Twitter container!",
		})
	})

	r.Run()
}
