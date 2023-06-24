package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", home)
	log.Println("Server was started.\nListening...")
	http.ListenAndServe("localhost:8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	log.Println("==================Start===================")
	defer log.Println("===================End====================")
	ctx := r.Context()
	log.Println("Request just started...")
	defer log.Println("Ended request anyway")
	select {
	case <-ctx.Done():
		log.Println("<Request was canceled!>")
	case <-time.After(time.Second * 10):
		log.Println("Request finished normally!")
		w.Write([]byte("Request finished normally!"))
	}
}
