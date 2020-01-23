package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"time"
)

type TestData struct {
	Field1 string
	Field2 string
	Filed3 int
	Time time.Time
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hi")
	})

	http.HandleFunc("/api/v1/test", func(w http.ResponseWriter, r *http.Request) {
		data :=TestData{
			"Field 1",
			"File 2",
			4,
			time.Now(),
		}

		dataJson, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(dataJson)

	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}

