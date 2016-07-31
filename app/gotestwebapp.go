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
	http.HandleFunc("/hello", helloHandler)
	log.Printf("starting up on port %v...", *port)
	err := http.ListenAndServe(":"+*port, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("called mainHandler")
	fmt.Fprintf(w, myapi.Greeting()+", I am working...")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("called mainHandler")
	name := r.URL.Query().Get("name")
	log.Printf("param=%v",name)
	fmt.Fprintf(w, myapi.Greet(name))
}
