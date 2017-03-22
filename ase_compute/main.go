package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	// "gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello from the compute container",
		})
	})

	r.Run()
}
