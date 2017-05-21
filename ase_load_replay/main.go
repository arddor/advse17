// ase_twitter

package main

import (
	"fmt"
	
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	
	"github.com/streadway/amqp"
)

var amqpChannel *amqp.Channel

func main() {
	
	// RabbitQ
	fmt.Println("Connecting to AMQP...")
	initializeAMQPConnection()
	
	// mongo
	fmt.Println("Connecting to mongodb...")
	session, err := mgo.Dial("mongodb://tweets:27017") // docker
	
	if err != nil {
		panic(err)
	}
	defer session.Close()
	
	collection := session.DB("test").C("trump_tweets")
	
	
	
	body := "Hi I like Hamburgers"
	timestamp := time.Now().String()
	
	err := amqpChannel.Publish(
			"",     // exchange
			"tweet", // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "text/plain",
				MessageId:    timestamp,
				Body:         []byte(body),
			})
	if err != nil {
		log.Println("Error while publishing tweet: ", err)
	}



}

func initializeAMQPConnection() {
	var err error
	var conn *amqp.Connection
	var ch *amqp.Channel

	connError := make(chan *amqp.Error)
	go func() {
		err := <-connError
		log.Println("reconnect: ", err)
		conn, ch = connectToMQ()
		//TODO: is this neccessary?
		//conn.NotifyClose(connError)
	}()

	conn, amqpChannel = connectToMQ()
	conn.NotifyClose(connError)

	defer conn.Close()

	if(amqpChannel == nil){
		amqpChannel, err = conn.Channel()
		log.Println("Failed to open a channel: ", err)
	}

	q, err := amqpChannel.QueueDeclare(
		"tweet", // name
		true,    // durable -> queue is not "lost" even when rabbitMQ crashes
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		log.Println("Error while creating a Queue: ", err)
	}
	log.Println("Declared queue " + q.Name)
}

func connectToMQ() (*amqp.Connection, *amqp.Channel) {
	for {
		var conn *amqp.Connection
		var err error
		conn, err = amqp.Dial("amqp://queue:5672")
		if err == nil {
			for {
				channel, err := conn.Channel()
				if err == nil{
					// TODO: are those next 2 lines neccessary?
					defer conn.Close()
					defer channel.Close()

					return conn, channel
				}
				log.Println("Error with channel creation: ", err)
			}
		}
		// else, reconnect after timeout
		time.Sleep(1000 * time.Millisecond)
		log.Println("Reconnect to AMQP...")
	}
}

