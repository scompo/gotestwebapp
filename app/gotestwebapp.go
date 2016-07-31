package main

import (
	"flag"
	"fmt"
	"github.com/scompo/gotestwebapp/myapi"
	"io"
	"log"
	"net/http"
	"os"
)

var port = flag.String("port", "8080", "port to listen to")

func main() {
	log.Println("Initializing...")
	flag.Parse()
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/upload", uploadHandler)
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

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	out, err := os.Create("/tmp/file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("File %v loaded succesfully\n", header.Filename)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("called mainHandler")
	name := r.URL.Query().Get("name")
	log.Printf("param=%v", name)
	fmt.Fprintf(w, myapi.Greet(name))
}
