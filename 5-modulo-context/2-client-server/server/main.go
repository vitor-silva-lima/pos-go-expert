package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("request received")
	defer log.Println("request finished")
	select {
	case <-ctx.Done():
		log.Println("request cancelled")
		return
	case <-time.After(5 * time.Second):
		log.Println("request processed")
		w.Write([]byte("request processed"))
	}
}
