package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func Plaintext(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello World!\n")
}

type JSONStruct struct {
	Array  []int          `json:"array"`
	Dict   map[string]int `json:"dict"`
	Int    int            `json:"int"`
	String string         `json:"string"`
	Double float64        `json:"double"`
	Null   interface{}    `json:"null"`
}

func JSON(w http.ResponseWriter, req *http.Request) {
	j := JSONStruct{Array: []int{1, 2, 3},
		Dict:   map[string]int{"one": 1, "two": 2, "three": 3},
		Int:    42,
		String: "test",
		Double: 3.14,
		Null:   nil}

	b, _ := json.MarshalIndent(j, "", " ")
	io.WriteString(w, string(b))
}

type User struct {
	ID    int    `db:"id" json:"id,omitempty"`
	Name  string `db:"name" json:"name,omitempty"`
	Email string `db:"email" json:"email,omitempty"`
}

// typical usage would keep or cache the open DB connection
var db, _ = sqlx.Open("sqlite3", "../database/test.sqlite")

func SQLiteFetch(w http.ResponseWriter, req *http.Request) {
	user := User{}
	rows, err := db.Queryx("select * from users order by random() limit 1")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.StructScan(&user)
		if err != nil {
			log.Fatal(err)
		}

		b, _ := json.MarshalIndent(user, "", " ")
		io.WriteString(w, string(b))
	}
}

// simpler version of SQLiteFetch, but no different performance wise
func SQLiteFetchAlt(w http.ResponseWriter, req *http.Request) {
	var user User
	err := db.QueryRowx("select * from users order by random() limit 1").StructScan(&user)
	if err != nil {
		log.Fatal(err)
	}

	b, _ := json.MarshalIndent(user, "", " ")
	io.WriteString(w, string(b))
}

var portNumber int

func main() {
	flag.IntVar(&portNumber, "port", 8300, "port number to listen on")
	flag.Parse()

	http.HandleFunc("/plaintext", Plaintext)
	http.HandleFunc("/json", JSON)
	http.HandleFunc("/sqlite-fetch", SQLiteFetch)
	http.HandleFunc("/sqlite-fetch2", SQLiteFetchAlt)

	log.Println("bench running on", fmt.Sprintf("%d", portNumber))

	err := http.ListenAndServe(fmt.Sprintf(":%d", portNumber), nil)
	if err != nil {
		log.Fatal(err)
	}
}
