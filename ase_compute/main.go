package main

import (
	"fmt"
	"net/http"

	"github.com/cdipaolo/sentiment"
	"github.com/gin-gonic/gin"

	"gopkg.in/mgo.v2"
)

var (
	model sentiment.Models
)

func sentimentCdipaolo(sentence string) uint8 {
	// fmt.Println("cdipaolo: ", model.SentimentAnalysis(sentence, sentiment.English).Score)
	return model.SentimentAnalysis(sentence, sentiment.English).Score
}

// Tweet Text and sentiment value of a tweet
type Tweet struct {
	Text      string
	Sentiment uint8
}

func main() {

	var err error

	// sentiment analysis
	fmt.Println("starting sentiment analysis")
	model, err = sentiment.Restore()
	if err != nil {
		panic(fmt.Sprintf("Could not restore model!\n\t%v\n", err))
	}

	// mongo
	fmt.Println("connecting to mongodb")
	session, err := mgo.Dial("mongodb://127.0.0.1:27017") // local
	if err != nil {
		panic(err)
	}
	defer session.Close()

	collection := session.DB("test").C("tweets")

	// result := Tweet{}
	// err = collection.Find(bson.M{"sentiment": 1}).One(&result)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Text:", result.Text)

	// gin
	fmt.Println("starting gin")
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello from the compute container",
		})
	})

	r.POST("/insert", func(c *gin.Context) {

		message := c.Query("message")
		sentimentValue := sentimentCdipaolo(message)
		err = collection.Insert(&Tweet{message, sentimentValue})
		if err != nil {
			c.JSON(http.StatusConflict, gin.H{
				"message": "Insert failed: " + message,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "Insert successful: " + message,
			})
		}
	})
	fmt.Println("running ...")
	r.Run()

}
