// ase_twitter

package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/arddor/advse17/lib_db"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/streadway/amqp"
)

// Incoming tweets need to be directly stored into the queue
// https://github.com/streadway/amqp

// upon startup needs to go get the currently stored terms
// https://www.rethinkdb.com/docs/
// then start listening for term updates

var (
	trackingParams []string
)

func addTrackingParam(param string) {
	k := 0
	for _, n := range trackingParams {
		if n != param { // filter
			trackingParams[k] = n
			k++
		}
	}

	trackingParams = append(trackingParams, param)
}

func paramAlreadyTracked(param string) bool {

	for _, n := range trackingParams {
		if n == param { // filter
			return true
		}
	}
	return false
}

func removeTrackingParam(param string) {
	// TODO: check if this works
	k := 0
	for _, n := range trackingParams {
		if n != param { // filter
			trackingParams[k] = n
			k++
		}
	}
	trackingParams = trackingParams[:k] // set slice len to remaining elements
}

func connectToMQ() *amqp.Connection {
	for {
		var conn *amqp.Connection
		var err error
		conn, err = amqp.Dial("amqp://ase_queue:5672")
		if err == nil {
			return conn
		}
		// else, reconnect after timeout
		time.Sleep(1000 * time.Millisecond)
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		fmt.Sprintf("%s: %s", msg, err)
		fmt.Print("FailOnError")
	}
}

func main() {

	// authenticate
	// TODO: change this so the keys are not in clear text
	consumerKey := "TheYSOyWqkVy5LS4AFj10LrXy"
	consumerSecret := "Qf9ovSx4aqFK9NkycjD2q1YYos5VhNVcNUFjyUjhDY8x3PWHoP"
	accessToken := "49389452-QjuTHd6wbDJUnRsD8gRbEPN076QVLlTVtHirbtgBa"
	accessSecret := "MlUhiDtWbYtMa1w3xLERmcATc6WVXYRr69xKGmnpslsWt"

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	//rabbitCloseError chan *amqp.Error
	//rabbitCloseError = make(chan *amqp.Error)

	var err error
	var conn *amqp.Connection

	connError := make(chan *amqp.Error)
	go func() {
		err := <-connError
		log.Println("reconnect: " + err.Error())
		conn = connectToMQ()
	}()

	// connect to RabbitMQ server
	conn = connectToMQ()
	failOnError(err, "Failed to connect to RabbitMQ")
	conn.NotifyClose(connError)

	// create a channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer conn.Close()

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
	defer ch.Close()

	var stream *twitter.Stream

	client := twitter.NewClient(httpClient)

	db.Initialize("ase_timeseries:28015")

	// Convenience Demux demultiplexed stream messages
	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {

		text := tweet.Text
		timestamp := tweet.CreatedAt
		//fmt.Println(text)
		// publish a message to the queue
		body := text
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "text/plain",
				MessageId:    timestamp,
				Body:         []byte(body),
			})
		failOnError(err, "Failed to publish a message")
	}

	var terms []db.Term
	terms, error := db.GetTerms(false)

	if error != nil {
		fmt.Println(error)
	}

	for _, term := range terms {
		addTrackingParam(term.Term)
	}

	params := &twitter.StreamFilterParams{
		Track:         trackingParams,
		StallWarnings: twitter.Bool(true),
	}
	stream, err = client.Streams.Filter(params)

	if err != nil {
		log.Fatal(err)
	}
	// Receive messages until stopped or stream quits
	// @marc: TODO: is here a go routine sensible or should I leave the go out?
	// Is there a need to implement a 'quit' handle in the routine?
	go demux.HandleChan(stream.Messages)

	// This gets called a lot due to inserts of sentiments
	db.OnChange(func(change map[string]*db.Term) {
		var tempTerm *db.Term
		var oldTerm *db.Term
		tempTerm = change["new_val"]
		oldTerm = change["old_val"]
		fmt.Println("Change: ")
		if oldTerm != tempTerm {
			// TODO: check conditions

			if tempTerm == nil {
				removeTrackingParam(oldTerm.Term)
				fmt.Println(oldTerm.Term + " deleted.")
			}

			if (oldTerm == nil && tempTerm != nil) || tempTerm != nil {
				if paramAlreadyTracked(tempTerm.Term) {
					return
				}

				addTrackingParam(tempTerm.Term)
				fmt.Println(tempTerm.Term + " added.")

			}

			stream.Stop()
			params := &twitter.StreamFilterParams{
				Track:         trackingParams,
				StallWarnings: twitter.Bool(true),
			}
			stream, err = client.Streams.Filter(params)
			// @marc: TODO: can this be handled in that way or do I have to somehow close the opened
			// go routine and open a new one?
			//go demux.HandleChan(stream.Messages)
			demux.HandleChan(stream.Messages)
			if err != nil {
				log.Fatal(err)
			}
		}
	})

	// Wait for SIGINT and SIGTERM (HIT CTRL-C)
	channel := make(chan os.Signal)
	signal.Notify(channel, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-channel)

	fmt.Println("Stopping Stream...")
	stream.Stop()

}
