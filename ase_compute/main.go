// ase_compute

package main

import (
	"fmt"

	"github.com/cdipaolo/sentiment"

	"gopkg.in/mgo.v2/bson"

	"ase_api/db" // I copied this to "C:\Users\B\go\src\ase_api" to work locally -> remove ir later
)

// Needs to process the queue
// upon startup needs to go into "polling" mode: continuously poll the queue (https://github.com/streadway/amqp)
// if the queue doesn't return a value, wait for a sensible amount (1 second?)

// if new tweet gets polled, process it - this means
// first: match the tweet to a term: basically rebuild twitters matching (https://dev.twitter.com/streaming/overview/request-parameters#track)
// 			throw away the tweet if no match exists
// second: generate the sentiment
// finally: store the sentiment to the term in the mongo db

// Regarding the terms from mongo:
// Each container requires a session to the db anyway
// on startup query the the terms and start polling every 10 seconds

// TODO: instead of polling the db will push term updated
// https://www.rethinkdb.com/docs/
// ideally each compute node will get notified from the DB with new terms

var (
	_terms             []db.Term
	_number_of_threads = 10
)

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

func initSentimentAnalysis() sentiment.Models {
	fmt.Println("Starting sentiment analysis")
	var model, err = sentiment.Restore()
	if err != nil {
		panic(fmt.Sprintf("Could not restore model!\n\t%v\n", err))
	}
	return model
}

func sentimentCdipaolo(sentence string, model sentiment.Models) uint8 {
	var score = model.SentimentAnalysis(sentence, sentiment.English).Score
	fmt.Println("Sentiment for '"+sentence+"':", score)
	return score
}

func initDB() {
	var err error

	db.Initialize("ase_timeseries:28015")

	_terms, err = db.GetTerms()
	if err != nil {
		panic(fmt.Sprintf("Could not retrieve terms!\n\t%v\n", err))
	}

	db.OnChange(func(change map[string]*db.Term) {
		var newTerm *db.Term
		var oldTerm *db.Term
		newTerm = change["new_val"]
		oldTerm = change["old_val"]
		if oldTerm == nil { // term was added
			_terms = append(_terms, *newTerm)
		} else if newTerm == nil { // term was removed
			for index, t := range _terms {
				if t.Term == oldTerm.Term {
					_terms[index] = _terms[len(_terms)-1] // Replace it with the last one.
					_terms = _terms[:len(_terms)-1]       // Chop off the last one
					break
				}
			}
		}
	})
}

// func aseGetPostFormArray(c *gin.Context) url.Values {
// 	req := c.Request
// 	req.ParseForm()
// 	req.ParseMultipartForm(32 << 20) // 32 MB
// 	return req.PostForm
// }

// func initGin(model sentiment.Models) {
// 	var err error

// 	fmt.Println("Starting gin")
// 	r := gin.Default()

// 	r.GET("/", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "Hello from the compute container",
// 		})
// 	})

// 	r.POST("/insert", func(c *gin.Context) {
// 		// URL like this: http://ase_compute:8080/insert or http://localhost:5000/insert from outside docker
// 		// Header like this: Content-Type: application/x-www-form-urlencoded
// 		// Body like this: Content-Type: timestamp1=I+like+cake
// 		fmt.Println(c.Query("term"))
// 		tweets := aseGetPostFormArray(c) // get tweets from queue later

// 		// find all terms
// 		// ### TODO: Do this only when new term was registered
// 		var allTerms []AseTerm
// 		err = collection.Find(nil).Select(bson.M{"term": 1}).All(&allTerms)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		// ###

// 		if tweets != nil && len(tweets) > 0 { // at least 1 tweet should be passed (arrays are supported)

// 			for key, value := range tweets {
// 				for _, term := range allTerms {
// 					if strings.Contains(strings.ToLower(value[0]), strings.ToLower(term.Term)) {
// 						fmt.Println("'" + value[0] + "' contains " + term.Term)
// 						pushToArray := bson.M{"$push": bson.M{"data": AseSentimentData{sentimentCdipaolo(value[0], model), key}}}
// 						err = collection.Update(bson.M{"_id": term.ID}, pushToArray)
// 						if err == nil {
// 							fmt.Println("Update on " + term.ID + " successful")
// 						} else {
// 							fmt.Println("Update on " + term.ID + " NOT successful")
// 						}
// 					}
// 				}
// 			}
// 		} else {
// 			fmt.Println("no tweet received")
// 		}
// 	})
// 	fmt.Println("Running ...")
// 	r.Run()
// }

func startWorker(model sentiment.Models, terms []db.Term) {

}

func main() {

	var model = initSentimentAnalysis()

	initDB()

	for i := 0; i < _number_of_threads; i++ {
		go startWorker(model, _terms)
	}

}
