package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// DefaultPort is the default port to use if once is not specified by the SERVER_PORT environment variable
const DefaultPort = "8888"

func getServerPort() string {
	port := os.Getenv("SERVER_PORT")
	if port != "" {
		return port
	}

	return DefaultPort
}

// EchoHandler echos back the request as a response
func Handler(writer http.ResponseWriter, request *http.Request) {
	buf, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(buf))
}

func main() {
	log.Println("starting server, listening on port " + getServerPort())

	http.HandleFunc("/", Handler)
	http.ListenAndServe(":"+getServerPort(), nil)
}
