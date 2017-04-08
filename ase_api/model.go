package main

import (
	"fmt"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type term struct {
	ID   bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Term string        `json:"term"`
	Data []sentiment   `json:"data"`
}

type sentiment struct {
	Timestamp time.Time `json:"time"`
	Sentiment int       `json:"sentiment"`
}

func (t *term) createTerm(db *mgo.Session) error {
	session := db.Clone()
	defer session.Close()

	t.ID = bson.NewObjectId()
	c := session.DB("ase").C("terms")
	err := c.Insert(&t)

	if err != nil {
		return err
	}
	return nil
}

func getTerms(db *mgo.Session) ([]term, error) {
	session := db.Clone()
	defer session.Close()

	c := session.DB("ase").C("terms")

	var terms []term
	err := c.Find(nil).All(&terms)
	if err != nil {
		return nil, err
	}
	fmt.Println("Results All:", terms)

	return terms, nil
}
