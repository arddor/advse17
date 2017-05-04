//ase_api

package main

import (
	"log"
	"net/http"

	"ase_api/db"

	"github.com/gin-gonic/gin"
	r "gopkg.in/gorethink/gorethink.v3"
)

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
	s.Router.GET("/", s.listTerms)
	// s.Router.POST("/", s.createTerm)
}

func (s *Server) Run(addr string) {
	log.Fatal(s.Router.Run(addr))
}

func (s *Server) listTerms(c *gin.Context) {
	terms, err := db.GetTerms()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Error",
		})
	}
	c.JSON(http.StatusOK, terms)
}

/*
func (s *Server) createTerm(c *gin.Context) {
	var term term
	err := c.Bind(&term)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if err := term.createTerm(s.DB); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid request payload")
		return
	}

	c.JSON(http.StatusCreated, term)
}
*/
