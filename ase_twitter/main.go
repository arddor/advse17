// ase_twitter

package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"sync"

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

type context struct {
	rabbitChannel *amqp.Channel
	rabbitQueue *amqp.Queue
}

var (
	trackingParams []string
)

var (
	rabbitConn *amqp.Connection
	rabbitCloseError chan *amqp.Error
	connectionContext *context
	_mutex        sync.Mutex
)

func connectToRabbitMQ(uri string) *amqp.Connection {
	for {
		conn, err := amqp.Dial(uri)
		
		if err == nil {
			return conn
		}
		
		log.Println(err)
		log.Printf("Trying to reconnect to RabbitMQ...")
		time.Sleep(5000 * time.Millisecond)
	}
}

func rabbitConnector(uri string) {
	var rabbitErr *amqp.Error
	
	for {
		rabbitErr = <-rabbitCloseError
		if rabbitErr != nil {
			log.Printf("Connecting to RabbitMQ...")
			rabbitConn = connectToRabbitMQ(uri)
			rabbitCloseError = make(chan *amqp.Error)
			rabbitConn.NotifyClose(rabbitCloseError)
			// run your setup process here
			ch, err := rabbitConn.Channel()

			failOnError(err, "Failed to open a channel")
			//defer ch.Close()
			
			q, err := ch.QueueDeclare(
							"tweet", // name
							true,   // durable
							false,   // delete when unused
							false,   // exclusive
							false,   // no-wait
							nil,     // arguments
			)
			failOnError(err, "Failed to declare a queue")
//			_mutex.Lock()
			connectionContext = &context{rabbitChannel: ch, rabbitQueue: &q}
//			_mutex.Unlock()
			fmt.Println("Declared queue " + connectionContext.rabbitQueue.Name)			
		}
	}

}

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

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		fmt.Sprintf("%s: %s", msg, err)
		fmt.Print("FailOnError")
	}
}

func main() {
	_mutex.Lock()
	connectionContext = &context{rabbitChannel: nil, rabbitQueue: nil}
	_mutex.Unlock()
	// authenticate
	// TODO: change this so the keys are not in clear text
	consumerKey := "TheYSOyWqkVy5LS4AFj10LrXy"
	consumerSecret := "Qf9ovSx4aqFK9NkycjD2q1YYos5VhNVcNUFjyUjhDY8x3PWHoP"
	accessToken := "49389452-QjuTHd6wbDJUnRsD8gRbEPN076QVLlTVtHirbtgBa"
	accessSecret := "MlUhiDtWbYtMa1w3xLERmcATc6WVXYRr69xKGmnpslsWt"

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// RabbitQ
	fmt.Println("Connecting to RabbitMQ...")
	
	// create the rabbitmq error channel
	rabbitCloseError = make(chan *amqp.Error)
	
	//TODO: run the callback in a separate thread?
	go rabbitConnector("amqp://queue:5672")
	
	// establish the rabbitmq connection by sending
	// an error and thus calling the error callback
	rabbitCloseError <- amqp.ErrClosed
	
	// create a channel
	//ch, err := rabbitConn.Channel()
	//failOnError(err, "Failed to open a channel")
	//defer ch.Close()
	
	//q, err := ch.QueueDeclare(
//		"tweet", // name
//		true,   // durable
//		false,   // delete when unused
//		false,   // exclusive
//		false,   // no-wait
//		nil,     // arguments
//	)
	
//	failOnError(err, "Failed to declare a queue")
//	fmt.Println("Declared queue " + q.Name)
//	connectionContext = &context{rabbitChannel: ch, rabbitQueue: &q} 
	
	var stream *twitter.Stream
	var err error
	for {
		if connectionContext.rabbitChannel != nil {
		log.Println("not nil anymore")
			break
		}
		log.Println("nil")
		time.Sleep(5000 * time.Millisecond)
	}
	client := twitter.NewClient(httpClient)

	db.Initialize("timeseries-db:28015")
	log.Println("Log message ")

	// Convenience Demux demultiplexed stream messages
	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		err = connectionContext.rabbitChannel.Publish(
		  "",		// exchange
		  connectionContext.rabbitQueue.Name,	// routing key
		  false,		// mandatory
		  false,		// immediate
		  amqp.Publishing {
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			MessageId:    tweet.CreatedAt,
			Body:         []byte(tweet.Text),
		  })
		//failOnError(err, "Failed to publish a message")
		log.Println(err)
		log.Println(tweet.Text)
		time.Sleep(2000 * time.Millisecond)
		//TODO: just throw tweet away?
	}

	var terms []db.Term
	terms, error := db.GetTerms(false)

	if error != nil {
		fmt.Println(error)
	}

	for _, term := range terms {
		addTrackingParam(term.Term)
		log.Println(term.Term)
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
	log.Println("Log message ")
	
	db.OnAddTerm(func(term db.Term) {

		fmt.Println("Change: ")

		addTrackingParam(term.Term)
		fmt.Println(term.Term + " added.")

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
	})

	// Wait for SIGINT and SIGTERM (HIT CTRL-C)
	channel := make(chan os.Signal)
	signal.Notify(channel, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-channel)

	fmt.Println("Stopping Stream...")
	stream.Stop()

}
