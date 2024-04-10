package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, req *http.Request) {
	log.Println("handle incoming:", req.RequestURI)
	n := rand.Int63n(500)
	time.Sleep(time.Duration(n) * time.Millisecond)
	d := map[string]any{"h": req.RequestURI, "n": n}
	m, err := json.Marshal(d)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
	}
	fmt.Fprintf(w, string(m))
}

func panic(http.ResponseWriter, *http.Request) {
	log.Panic("panic on server")
}

func main() {
	s := http.NewServeMux()
	s.HandleFunc("/kill", panic)
	s.HandleFunc("/", handler)
	port := "8090"
	log.Println("Starting app on port", port)
	log.Fatal(http.ListenAndServe(":"+port, s))
}
