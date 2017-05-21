// ase_api

package main

import (
	"log"
	"net/http"
	"time"

	"github.com/arddor/advse17/lib_db"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	r "gopkg.in/gorethink/gorethink.v3"
	"github.com/streadway/amqp"
)

// TODO: Add check of origin again
var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

func main() {
	s := Server{}
	s.Initialize("timeseries-db:28015")
	s.Run(":8000")
}

type Server struct {
	Router *gin.Engine
	DB     *r.Session
}

func (s *Server) Initialize(addr string) {
	s.DB = db.Initialize(addr)

	s.Router = gin.Default()
	s.initializeRoutes()
	
	initializeAMQPConnection()
}

func (s *Server) initializeRoutes() {
	s.Router.GET("/terms", s.listTerms)
	s.Router.POST("/terms", s.createTerm)
	s.Router.GET("/terms/:id", s.getTerm)
	s.Router.GET("/echo", func(c *gin.Context) {
		wsHandler(c.Writer, c.Request)
	})
}

func (s *Server) Run(addr string) {
	log.Fatal(s.Router.Run(addr))
}

func (s *Server) listTerms(c *gin.Context) {
	terms, err := db.GetTerms(false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error",
		})
	}
	c.JSON(http.StatusOK, terms)
}

func (s *Server) createTerm(c *gin.Context) {
	var param db.Term
	
	// TODO: needs check if already exists

	c.BindJSON(&param)

	if param.Term == "" {
		c.JSON(http.StatusBadRequest, "Term was empty ")
		return
	}

	term, err := db.CreateTerm(param.Term)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, term)
}

 func (s *Server) createQueueItem(c *gin.Context) {
     // TODO: test
	var item string
	
	c.BindJSON(&item)
	
	if item == "" {
		c.JSON(http.StatusBadRequest, "Item was empty ")
		return
	}
	
	body := item
	// TODO: if you want to have a different or specific format, you need to use the .Format("...") function
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
	c.JSON(http.StatusCreated, item)
}

func (s *Server) getTerm(c *gin.Context) {
	id := c.Param("id")
	term, err := db.GetTerm(id, true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, term)
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	err = c.WriteMessage(websocket.TextMessage, []byte("Test"))
	if err != nil {
		log.Println("write:", err)
	}
}
