package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		select {
		default:
		case <-r.Context().Done():
			return
		}
		<-time.After(5 * time.Second)
	})
	log.Print("Starting web server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
