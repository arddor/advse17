//ase_api

package main

import (
	"ase_api/db"
	"log"
	"net/http"

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
	s.Router.GET("/terms", s.listTerms)
	s.Router.POST("/terms", s.createTerm)
	s.Router.GET("/terms/:id", s.getTerm)
}

func (s *Server) Run(addr string) {
	log.Fatal(s.Router.Run(addr))
}

func (s *Server) listTerms(c *gin.Context) {
	terms, err := db.GetTerms()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error",
		})
	}
	c.JSON(http.StatusOK, terms)
}

func (s *Server) createTerm(c *gin.Context) {
	var param db.Term
	
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
	term, err := db.GetTerm(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, term)
}

type Test struct {Term string `json:"term" gorethink:"term"` }

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
