// ase_compute

package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	hc "cirello.io/HumorChecker"

	"github.com/arddor/advse17/lib_db"

	"sync"

	"github.com/streadway/amqp"
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
	_terms        []db.Term
	_mutex        sync.Mutex
	_maxSentiment = 6
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func printLog(prefix string, msg string) {
	fmt.Println("[" + prefix + "] " + msg)
}

func clipping(score int, max int) int {
	if score < -max {
		score = -max
	} else if score > max {
		score = max
	}
	return score
}

func computeSentiment(sentence string) float32 {
	var rawScore = clipping(hc.Analyze(sentence).Score, _maxSentiment)

	var score = float32(_maxSentiment+rawScore) / float32(2*_maxSentiment)
	printLog("Sentiment", "Sentiment Score:")
	fmt.Println(score)
	return score
}

func initDB() {
	printLog("DB", "Starting ...")

	db.Initialize("ase_timeseries:28015")

	_terms, _ = db.GetTerms(false)
	if _terms == nil {
		printLog("DB", "No terms registered yet")
	}
	go func() {
		db.OnChangeNoData(func(change map[string]*db.Term) {
			var newTerm *db.Term
			var oldTerm *db.Term
			newTerm = change["new_val"]
			oldTerm = change["old_val"]

			_mutex.Lock()
			if oldTerm == nil { // term was added
				_terms = append(_terms, *newTerm)
				printLog("DB", "Term added: "+newTerm.Term)
			} else if newTerm == nil { // term was removed
				for index, t := range _terms {
					if t.Term == oldTerm.Term {
						_terms[index] = _terms[len(_terms)-1] // Replace it with the last one.
						_terms = _terms[:len(_terms)-1]       // Chop off the last one
						printLog("DB", "Term deleted: "+oldTerm.Term)
						break
					}
				}
			}
			_mutex.Unlock()
		})
	}()
}

func processTweet(timestamp string, tweet string) bool {
	_mutex.Lock()
	for _, term := range _terms {
		if strings.Contains(strings.ToLower(tweet), strings.ToLower(term.Term)) {
			_mutex.Unlock()
			printLog("ProcessTweet", "Tweet contains "+term.Term)
			fmt.Print("'" + timestamp + "' converted to: ")
			// layout := "2006-01-02T15:04:05.000Z" // Example
			layout := "Mon Jan 02 15:04:05 +0000 2006"
			t, err := time.Parse(layout, timestamp)
			fmt.Print("'" + t.String() + "'\n")
			if err != nil {
				printLog("ProcessTweet", "Could not convert timestamp!")
				return false
			}
			var sentimentData = db.Sentiment{Timestamp: t, Sentiment: computeSentiment(tweet)}
			err = db.AddSentiment(term.ID, sentimentData)
			if err != nil {
				printLog("ProcessTweet", "Could not add sentiment to DB!")
				return false
			}
			return true
		}
	}
	_mutex.Unlock()
	printLog("ProcessTweet", "Tweet does not contain a term")
	return true
}

func startWorker() {
	var conn *amqp.Connection
	printLog("Worker", "Starting ...")
	for {
		var err error
		conn, err = amqp.Dial("amqp://ase_queue:5672")
		if err == nil {
			break
		} else {
			time.Sleep(1000 * time.Millisecond)
		}
	}
	printLog("Worker", "Queue connected")
	defer conn.Close()

	// create a channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// declare a queue for us to send to
	q, err := ch.QueueDeclare(
		"tweet", // name
		true,    // durable -> queue is not "lost" even when rabbitMQ crashes
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			printLog("Worker", "Received a tweet: "+string(d.Body))
			if processTweet(d.MessageId, string(d.Body)) {
				d.Ack(false)
			} else {
				d.Nack(false, true)
			}
		}
	}()

	printLog("Worker", " Waiting for messages. To exit press CTRL+C")
	<-forever
}

func main() {
	time.Sleep(3000 * time.Millisecond)

	initDB()

	startWorker()

}
