//LoadGeneration/ase_load

package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"net/http"
	"net/url"
	"time"
	
	"github.com/dghubble/oauth1"
	"github.com/gin-gonic/gin"
	"github.com/dghubble/go-twitter/twitter"
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
trackingParams = append(trackingParams, param)
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

//TODO: docker dependency
// use docker-compose depends_on
time.Sleep(3000 * time.Millisecond)

var stream *twitter.Stream
var err error
client := twitter.NewClient(httpClient)
// gin listening to: intern 8080, extern 5000
r := gin.Default()

r.PUT("/terms/:term", func(c *gin.Context) {
 
// Convenience Demux demultiplexed stream messages
	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
	// TODO: we need to send the actual term to the sentiment analysis
		text := tweet.Text
		timestamp := tweet.CreatedAt
		link := "http://ase_compute:8080/insert?term=test"
		resp, err := http.PostForm(link, url.Values{timestamp: {text}})
		if err != nil {
		log.Fatal(err)
		}
		fmt.Println(resp)
		defer resp.Body.Close()
	}
	demux.DM = func(dm *twitter.DirectMessage) {
		fmt.Println(dm.SenderID)
	}
	demux.Event = func(event *twitter.Event) {
		fmt.Printf("%#v\n", event)
	}


	fmt.Println("Starting Stream...")
	param := c.Param("term")
	fmt.Println("Found new param: " + param)
	// TODO: does this do anything?

	addTrackingParam(param)
	if len(trackingParams) > 1 {
		stream.Stop()
	}
	
	params := &twitter.StreamFilterParams{
		Track: trackingParams,
		StallWarnings: twitter.Bool(true),
		}
	stream, err = client.Streams.Filter(params)

	if err != nil {
		log.Fatal(err)
	}
	// Receive messages until stopped or stream quits
	go demux.HandleChan(stream.Messages)


	// Wait for SIGINT and SIGTERM (HIT CTRL-C)
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)

	fmt.Println("Stopping Stream...")
	stream.Stop()
	})
	r.Run()
}
