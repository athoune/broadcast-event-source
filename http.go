package main

import (
	"log"
	"net/http"
)

func blunderServer(w http.ResponseWriter, req *http.Request) {
	h := w.Header()
	h.Set("Content-Type", "text/event-stream")
	h.Set("Transfer-Encoding", "chunked")
	h.Set("Cache-Control", "no-cache")
	h.Set("Connection", "keep-alive")
	h.Set("X-Accel-Buffering", "no")
	w.WriteHeader(200)

}

func startHttp() {
	http.HandleFunc("/blunder", blunderServer)
	log.Printf("About to start http://localhost:8000")
	err := http.ListenAndServe("localhost:8000", nil)
	if err != nil {
		panic(err)
	}

}
