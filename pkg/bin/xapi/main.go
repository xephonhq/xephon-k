package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// This demostrate the API from spec

func main() {
	log.Print("Let's start API server")
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// The "/" pattern matches everything, need to check if it's really the root request
		if req.URL.Path != "/" {
			http.NotFound(w, req)
			return
		}
		// TODO: should return API doc etc.
		fmt.Fprintf(w, "welcom to home page!")
	})
	mux.HandleFunc("/info/version", func(w http.ResponseWriter, req *http.Request) {
		v := map[string]string{"version": "0.0.1"}
		b, err := json.Marshal(v)
		if err != nil {
			log.Fatal(err)
		}
		// FIXED: this would result in an array being printed out in the browser
		// fmt.Fprint(w, b)
		// TODO: is there a way to use the bytes directly without converting it to string?
		fmt.Fprint(w, string(b))
	})
	// TODO: the real logic for writing into Cassandra
	mux.HandleFunc("/w", func(w http.ResponseWriter, req *http.Request) {
		// TODO: check if the input body is valid
		if req.Method != "POST" {
			fmt.Fprintf(w, "you should use post!")
		} else {
			fmt.Fprintf(w, "nice post, but I don't know what to do /w\\")
		}
	})
	// TODO: the real logic for reading from Cassandra
	mux.HandleFunc("/q", func(w http.ResponseWriter, req *http.Request) {
		// TODO: check if the input body is valid
		if req.Method != "POST" {
			fmt.Fprintf(w, "you should use post!")
		} else {
			fmt.Fprintf(w, "nice post, but I don't know what to do /w\\")
		}
	})
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
