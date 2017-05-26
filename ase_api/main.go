// ase_api

package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/arddor/advse17/lib_db"
	"github.com/gin-contrib/static"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	r "gopkg.in/gorethink/gorethink.v3"
)

// TODO: Add check of origin again
// var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
var upgrader = &websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}

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
}

func (s *Server) initializeRoutes() {
	s.Router.Use(static.Serve("/", static.LocalFile("/public", true)))
	api := s.Router.Group("/api")
	{
		api.GET("/terms", s.listTerms)
		api.POST("/terms", s.createTerm)
		api.GET("/terms/:id", s.getTerm)
	}
	s.Router.GET("/ws/changes/:id", changesHandler())
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

func (s *Server) getTerm(c *gin.Context) {
	id := c.Param("id")
	seconds, _ := strconv.Atoi(c.Query("seconds"))

	// set 1h if missing
	if seconds <= 0 {
		seconds = 3600
	}

	term, err := db.GetTerm(id, seconds)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, term)
}

func changesHandler() gin.HandlerFunc {
	h := newHub()
	go h.run()

	// send to hub for broadcast
	go func() {
		for {
			db.OnAddSentiment(func(value interface{}) {
				h.broadcast <- value
			})
		}
	}()
	return func(c *gin.Context) {
		id := c.Param("id")
		wsHandler(id, h, c.Writer, c.Request)
	}
}

func wsHandler(id string, h hub, w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	c := &connection{id: id, send: make(chan interface{}, 256), ws: ws}
	h.register <- c
	defer func() { h.unregister <- c }()
	go c.writer()
	c.reader()
}

type connection struct {
	id   string
	ws   *websocket.Conn
	send chan interface{}
}

func (c *connection) reader() {
	for {
		if _, _, err := c.ws.ReadMessage(); err != nil {
			break
		}
	}
	c.ws.Close()
}

func (c *connection) writer() {
	for change := range c.send {
		message := map[string]interface{}(change.(map[string]interface{}))
		if message["id"] != c.id {
			continue
		}
		if err := c.ws.WriteJSON(change); err != nil {
			break
		}
	}
	c.ws.Close()
}

type hub struct {
	// Registered connections.
	connections map[*connection]bool

	// Inbound messages from the connections.
	broadcast chan interface{}

	// Register requests from the connections.
	register chan *connection

	// Unregister requests from connections.
	unregister chan *connection
}

func newHub() hub {
	return hub{
		broadcast:   make(chan interface{}),
		register:    make(chan *connection),
		unregister:  make(chan *connection),
		connections: make(map[*connection]bool),
	}
}

func (h *hub) run() {
	for {
		select {
		case c := <-h.register:
			h.connections[c] = true
		case c := <-h.unregister:
			if _, ok := h.connections[c]; ok {
				delete(h.connections, c)
				close(c.send)
			}
		case m := <-h.broadcast:
			for c := range h.connections {
				select {
				case c.send <- m:
				default:
					delete(h.connections, c)
					close(c.send)
				}
			}
		}
	}
}
