package db

import (
	"log"
	"time"

	r "gopkg.in/gorethink/gorethink.v3"
)

var session *r.Session

type Term struct {
	ID      string      `json:"id" gorethink:"id,omitempty"`
	Term    string      `json:"term"`
	Data    []Sentiment `json:"data"`
	Created time.Time
}

type Sentiment struct {
	Timestamp time.Time `json:"time"`
	Sentiment int       `json:"sentiment"`
}

func Initialize(addr string) *r.Session {
	var err error
	session, err = r.Connect(r.ConnectOpts{
		Address:  addr,
		Database: "term",
	})
	if err != nil {
		log.Fatalln(err.Error())
		panic("Connection could not be established")
	}
	return session
}

func GetTerms() ([]Term, error) {
	var terms []Term

	res, err := r.Table("items").Run(session)
	if err != nil {
		return nil, err
	}

	err = res.All(&terms)
	if err != nil {
		return nil, err
	}

	return terms, nil
}

func CreateTerm(term string) error {
	_, err := r.Table("items").Insert(Term{Term: term}).RunWrite(session)
	if err != nil {
		return err
	}
	return nil
}

func OnChange(fn func(value map[string]*Term)) {
	res, err := r.Table("items").Changes().Run(session)

	var value map[string]*Term

	if err != nil {
		log.Fatalln(err)
	}

	for res.Next(&value) {
		fn(value)
	}
}
