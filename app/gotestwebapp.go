package main

import (
	"flag"
	"fmt"
	"github.com/scompo/gotestwebapp/myapi"
	"log"
	"net/http"
)

var port = flag.String("port", "8080", "port to listen to")

func main() {
	log.Println("Initializing...")
	flag.Parse()
	http.HandleFunc("/", mainHandler)
	log.Printf("starting up on port %v...", *port)
	err := http.ListenAndServe(":"+*port, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, myapi.Greeting()+", I am working...")
}
