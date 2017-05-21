// ase_twitter

package main

import (
	"fmt"
	"time"
	
	// https://labix.org/mgo
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	
	"github.com/streadway/amqp"
)

type Tweet struct {
	Id		bson.ObjectId `json:"id" bson:"_id"`
	Text		string        `json:"text" bson:"text"`
}

func failOnError(err error, msg string) {
	if err != nil {
		fmt.Sprintf("%s: %s", msg, err)
	}
}

func main() {
	
	// mongo
	fmt.Println("Connecting to mongodb...")
	session, err := mgo.Dial("mongodb://tweets:27017")
	
	if err != nil {
		panic(err)
	}
	defer session.Close()
	
	coll := session.DB("test").C("trump_tweets")
	
	// RabbitQ
	fmt.Println("Connecting to RabbitQ...")
	
	var conn *amqp.Connection
	var connectError error
	
	for {
		conn, connectError = amqp.Dial("amqp://queue:5672")
		if connectError == nil {
			break
		} else {
			time.Sleep(1000 * time.Millisecond)
		}
	}

	fmt.Println("Queue connected")
	defer conn.Close()
	
	// create a channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
	
	q, err := ch.QueueDeclare(
		"tweet", // name
		true,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
	fmt.Println("Declared queue " + q.Name)
	
	// loop through all items
	
	result := Tweet{}
	
	for {
		iter := coll.Find(nil).Batch(10000).Iter()
		for iter.Next(&result) {		
			err = ch.Publish(
			  "",		// exchange
			  "tweet",	// routing key
			  false,		// mandatory
			  false,		// immediate
			  amqp.Publishing {
				DeliveryMode: amqp.Persistent,
				ContentType:  "text/plain",
				MessageId:    time.Now().Format("Mon Jan 02 15:04:05 +0000 2006"),
				Body:         []byte(result.Text),
			  })
			failOnError(err, "Failed to publish a message")
		}
		
		if err = iter.Close(); err != nil {
			failOnError(err, "Failed to close iterator")
		}
	}
	
	fmt.Println("Exiting replay...")
}
