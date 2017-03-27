package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	// "github.com/dghubble/go-twitter/twitter"
)

// So we need to first authenticate ourselfs: https://github.com/ChimeraCoder/anaconda#authentication
// 
// We then need to open a public stream https://dev.twitter.com/streaming/public
// 
// and use the POST endpoint https://dev.twitter.com/streaming/reference/post/statuses/filter
// 
// then pass our terms we'd like to observe with the "track" parameter https://dev.twitter.com/streaming/reference/post/statuses/filter
// 
// Ideally we would want ONE STREAM PER TERM we are tracking, that would make our life a lot easier.
// 
// Under that assumption this code here would just have to provide someway to receive a term and then start tracking it
// (aka authorize with twitter and start the stream)
// The received tweets should be processed and we should have an array of Strings with the content of the strings.

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello from the Twitter container!",
		})
	})

	r.Run()
}
