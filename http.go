package main

import (
	"fmt"
	"log"
	"net/http"
)

func blunderServer(w http.ResponseWriter, req *http.Request) {
	h := w.Header()
	h.Set("Content-Type", "text/event-stream")
	h.Set("Transfer-Encoding", "chunked")
	h.Set("Cache-Control", "no-cache")
	h.Set("Connection", "keep-alive")
	h.Set("X-Accel-Buffering", "no") // For Nginx
	w.WriteHeader(200)
	channel := Subscribe()
	defer func() {
		log.Println("Closing http connection closes is channel.")
		/*FIMXE finding how to close channel when curl make a ctrl-C*/
		channel.Leave()
	}()
	f, f_ok := w.(http.Flusher)
	if f_ok {
		f.Flush()
	}
	for {
		m := <-channel.channel
		msg := fmt.Sprintf("data: %s\r\n", m)
		w.Write([]byte(msg))
		if f_ok {
			f.Flush()
		}
	}
}

func startHttp() {
	http.HandleFunc("/blunder", blunderServer)
	http.Handle("/", http.FileServer(http.Dir("static")))
	log.Printf("About to start http://localhost:8000")
	err := http.ListenAndServe("localhost:8000", nil)
	if err != nil {
		panic(err)
	}
}
