// ase_api/db

package db

import (
	"fmt"
	"log"
	"strings"
	"time"

	r "gopkg.in/gorethink/gorethink.v3"
)

var session *r.Session

type Term struct {
	ID      string      `json:"id" gorethink:"id,omitempty"`
	Term    string      `json:"term" gorethink:"term"`
	Data    []Sentiment `json:"data, omitempty" gorethink:"data"`
	Created time.Time   `json:"created" gorethink:"created"`
}

type Sentiment struct {
	Timestamp time.Time `json:"time" gorethink:"timestamp"`
	Sentiment float32   `json:"sentiment" gorethink:"sentiment"`
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

	// Create Database if not exists
	r.DBList().Contains("term").Do(func(exists r.Term) r.Term {
		return r.Branch(exists, "do nothing", r.DBCreate("term"))
	}).Exec(session)
	// Create Table if not exists
	r.DB("term").TableList().Contains("items").Do(func(exists r.Term) r.Term {
		return r.Branch(exists, "do nothing", r.DB("term").TableCreate("items"))
	}).Exec(session)

	return session
}

func GetTerms(includeSentimentData bool) ([]Term, error) {
	var terms []Term
	var res *r.Cursor
	var err error

	if includeSentimentData {
		res, err = r.Table("items").Run(session)
	} else {
		res, err = r.Table("items").Without("data").Default([]Term{}).Run(session)
	}

	if err != nil {
		return nil, err
	}

	err = res.All(&terms)
	if err != nil {
		return nil, err
	}

	return terms, nil
}

func GetTerm(id string, seconds int) (*Term, error) {
	var cursor *r.Cursor
	var err error

	if seconds > 0 {
		cursor, err = r.Table("items").Get(id).Merge(map[string]interface{}{
			"data": r.Row.Field("data").Filter(func(item r.Term) r.Term {
				return item.Field("timestamp").Gt(r.Now().Sub(seconds))
			}),
		}).Run(session)
	} else {
		cursor, err = r.Table("items").Get(id).Without("data").Run(session)
	}

	if err != nil {
		return nil, err
	}

	var term Term
	cursor.One(&term)
	cursor.Close()

	return &term, nil
}

func CreateTerm(term string) (*Term, error) {
	obj := Term{
		Term:    term,
		Data:    []Sentiment{},
		Created: time.Now(),
	}

	res, err := r.Table("items").Filter(func(doc r.Term) r.Term {
		return doc.Field("term").Downcase().Eq(strings.ToLower(term))
	}).Count().Do(func(result r.Term) r.Term {
		return r.Branch(
			result.Eq(0),
			r.Table("items").Insert(obj),
			"exists",
		)
	}).RunWrite(session)
	if err != nil {
		fmt.Println("Some error ...")
		// log.Fatal(err)
		return nil, err
	}

	obj.ID = res.GeneratedKeys[0]

	return &obj, nil
}

// AddSentiment pushes a sentiment into the term with the id
func AddSentiment(id string, sentiment Sentiment) error {
	_, err := r.Table("items").Get(id).
		Update(map[string]interface{}{"data": r.Row.Field("data").Append(sentiment)}).
		RunWrite(session)

	return err
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

func OnAddSentiment(fn func(value interface{})) {
	res, err := r.Table("items").Pluck("data", "id").Changes().Map(func(doc r.Term) map[string]interface{} {
		return map[string]interface{}{
			"id":   doc.Field("new_val").Field("id"),
			"data": doc.Field("new_val").Field("data").Nth(-1),
		}
	}).Run(session)

	var value interface{}

	if err != nil {
		log.Fatalln(err)
	}

	for res.Next(&value) {
		fn(value)
	}
}

func OnAddTerm(fn func(term Term)) {
	res, err := r.Table("items").Pluck("term").Changes().Map(func(doc r.Term) interface{} {
		return doc.Field("new_val")
	}).Run(session)

	if err != nil {
		log.Fatalln(err)
	}

	var value Term

	for res.Next(&value) {
		fn(value)
	}
}

func OnChangeNoData(fn func(value map[string]*Term)) {
	res, err := r.Table("items").Without("data").Changes().Run(session)

	var value map[string]*Term

	if err != nil {
		log.Fatalln(err)
	}

	for res.Next(&value) {
		fn(value)
	}
}
