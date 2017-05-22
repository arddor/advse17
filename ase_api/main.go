// ase_api

package main

import (
	"log"
	"net/http"

	"github.com/arddor/advse17/lib_db"
	"github.com/gin-contrib/static"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	r "gopkg.in/gorethink/gorethink.v3"
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
}

func (s *Server) initializeRoutes() {
	s.Router.Use(static.Serve("/", static.LocalFile("public", true)))
	api := s.Router.Group("/api")
	{
		api.GET("/terms", s.listTerms)
		api.POST("/terms", s.createTerm)
		api.GET("/terms/:id", s.getTerm)
	}
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
