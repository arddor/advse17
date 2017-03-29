package main

import (
	"fmt"
	"net/http"

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
	// session, err := mgo.Dial("mongodb://127.0.0.1:27017") // local
	session, err := mgo.Dial("mongodb://ase_timeseries:27017") // docker

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
		// URL like this: http://localhost:5000/insert?term=Novartis
		// Header like this: Content-Type: application/x-www-form-urlencoded
		// Body like this: Content-Type: timestamp1=I+like+cake
		term := c.Query("term")
		tweets := aseGetPostFormArray(c)

		if tweets != nil && len(tweets) > 0 { // at least 1 tweet should be passed (arrays are supported)

			var sentimentData []AseSentimentData
			for key, value := range tweets {
				sentimentData = append(sentimentData, AseSentimentData{sentimentCdipaolo(value[0]), key})
			}

			termExists := AseTerm{}
			err = collection.Find(bson.M{"term": term}).Limit(1).One(&termExists) // figure out if term exists in database
			fmt.Println("Error after searching for term: ", err)

			if err == nil { // found a match, append to data
				sentimentData = append(termExists.Data, sentimentData...)
				pushToArray := bson.M{"$set": bson.M{"data": sentimentData}}
				fmt.Println(termExists.ID, pushToArray)
				err = collection.Update(bson.M{"_id": termExists.ID}, pushToArray)
				if err == nil {
					fmt.Println("Update on " + termExists.ID + " successful")
				}
			} else if err.Error() == "not found" { // term not yet in db
				err = collection.Insert(&AseTerm{bson.NewObjectId(), term, sentimentData})
				if err == nil {
					fmt.Println("Insert successful")
				}
			}

			// show errors
			if err != nil {
				c.JSON(http.StatusConflict, gin.H{
					"message": "Storing sentiments for " + term + " failed",
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"message": "Storing sentiments for " + term + " successful!",
				})
			}
		} else {
			fmt.Println("no tweet received")
		}
	})
	fmt.Println("Running ...")
	r.Run()

}
