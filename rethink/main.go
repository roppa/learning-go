package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	r "gopkg.in/rethinkdb/rethinkdb-go.v5"
)

var session *r.Session

func init() {
	var err error
	session, err = r.Connect(r.ConnectOpts{
		Address:    "localhost:28015",
		Database:   "example",
		InitialCap: 10,
		MaxOpen:    10,
	})
	if err != nil {
		log.Fatalln(err)
	}
}

func getFooData(write http.ResponseWriter, read *http.Request) {
	var all []interface{}
	r.DB("example").Table("foo").ReadAll(&all, session)
	data, err := json.Marshal(all)
	if err != nil {
		log.Printf("Error with json martial: %+v", err)
	}

	write.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(write, string(data))
}

// func watchFooData() {
// 	opts := r.ChangesOpts{IncludeInitial: true, IncludeStates: true}
//   cursor, err := r.DB("example").Table("foo").Get(123).Changes(opts).Run(session)
// }

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.HandleFunc("/data", getFooData)

	log.Fatal(http.ListenAndServe(":8081", nil))
}
