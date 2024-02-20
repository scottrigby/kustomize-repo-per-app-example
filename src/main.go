package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handler)

	go func() {
		log.Println("Starting server")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal(err)
		}
	}()

	done := make(chan bool)
	<-done
}

func handler(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("HELLO")
	if name == "" {
		name = "world"
	}
	log.Printf("Serving Hello %s\n", name)
	w.Write([]byte(fmt.Sprintf("Hello %s\n", name)))
}
