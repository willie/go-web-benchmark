package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
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

	b, _ := json.MarshalIndent(j, "", "\t") // since the output requested pretty formatting
	io.WriteString(w, string(b))

	json.NewEncoder(w).Encode(j)
}

var portNumber int

func main() {
	flag.IntVar(&portNumber, "port", 8300, "port number to listen on")
	flag.Parse()

	flag.VisitAll(func(f *flag.Flag) {
		log.Println(f.Name, f.Value)
	})
	log.Println("----")

	http.HandleFunc("/plaintext", Plaintext)
	http.HandleFunc("/json", JSON)

	err := http.ListenAndServe(fmt.Sprintf(":%d", portNumber), nil)
	if err != nil {
		log.Fatal(err)
	}
}
