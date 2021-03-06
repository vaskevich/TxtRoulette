package main

import (
	"fmt"
	"github.com/ovaskevich/TxtRoulette/server"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// Read the args.
	if len(os.Args) != 2 {
		log.Fatal("usage: server.go port")
	}
	port := ":" + os.Args[1]

	// Make sure environment variables are set.
	if len(os.Getenv("TWILIO_APIUSR")) == 0 || len(os.Getenv("TWILIO_APIKEY")) == 0 {
		log.Fatal("Please set your TWILIO_APIUSR and TWILIO_APIKEY environment variables.")
	}

	// Start the server.
	fmt.Printf("Starting TxtRoulette server on port %s...\n", port)
	http.HandleFunc("/receive/", server.Receive)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Welcome to TxtRoulette! Text CONNECT to 320-839-8785 to join.")
	})
	log.Fatal(http.ListenAndServe(port, nil))
}
