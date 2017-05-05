package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/cdipaolo/sentiment"
	"github.com/gin-gonic/gin"

	"net/url"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	model sentiment.Models
)

func sentimentCdipaolo(sentence string) uint8 {
	fmt.Println("Sentiment for '"+sentence+"':", model.SentimentAnalysis(sentence, sentiment.English).Score)
	return model.SentimentAnalysis(sentence, sentiment.English).Score
}

// AseSentimentData sentiment data including timestamp
type AseSentimentData struct {
	Sentiment uint8
	Timestamp string
}

// AseTerm term with associated data
type AseTerm struct {
	ID   bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Term string
	Data []AseSentimentData
}

func aseGetPostFormArray(c *gin.Context) url.Values {
	req := c.Request
	req.ParseForm()
	req.ParseMultipartForm(32 << 20) // 32 MB
	return req.PostForm
}

func main() {

	var err error

	// sentiment analysis
	fmt.Println("Starting sentiment analysis")
	model, err = sentiment.Restore()
	if err != nil {
		panic(fmt.Sprintf("Could not restore model!\n\t%v\n", err))
	}

	// mongo
	fmt.Println("Connecting to mongodb")
	session, err := mgo.Dial("mongodb://127.0.0.1:27017") // local
	// session, err := mgo.Dial("mongodb://ase_timeseries:27017") // docker

	if err != nil {
		panic(err)
	}
	defer session.Close()

	collection := session.DB("test").C("terms")

	// gin
	fmt.Println("Starting gin")
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello from the compute container",
		})
	})

	r.POST("/insert", func(c *gin.Context) {
		// URL like this: http://ase_compute:8080/insert or http://localhost:5000/insert from outside docker
		// Header like this: Content-Type: application/x-www-form-urlencoded
		// Body like this: Content-Type: timestamp1=I+like+cake
		fmt.Println(c.Query("term"))
		tweets := aseGetPostFormArray(c) // get tweets from queue later

		// find all terms
		// ### TODO: Do this only when new term was registered
		var allTerms []AseTerm
		err = collection.Find(nil).Select(bson.M{"term": 1}).All(&allTerms)
		if err != nil {
			fmt.Println(err)
		}
		// ###

		if tweets != nil && len(tweets) > 0 { // at least 1 tweet should be passed (arrays are supported)

			for key, value := range tweets {
				for _, term := range allTerms {
					if strings.Contains(strings.ToLower(value[0]), strings.ToLower(term.Term)) {
						fmt.Println("'" + value[0] + "' contains " + term.Term)
						pushToArray := bson.M{"$push": bson.M{"data": AseSentimentData{sentimentCdipaolo(value[0]), key}}}
						err = collection.Update(bson.M{"_id": term.ID}, pushToArray)
						if err == nil {
							fmt.Println("Update on " + term.ID + " successful")
						} else {
							fmt.Println("Update on " + term.ID + " NOT successful")
						}
					}
				}
			}
		} else {
			fmt.Println("no tweet received")
		}
	})
	fmt.Println("Running ...")
	r.Run()

}
