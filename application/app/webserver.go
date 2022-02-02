package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

// var shared between function calls
var hostName string

func init() {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	hostName = name
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Method: " + r.Method)

	query := r.URL.Query()
	user := query.Get("user")
	time := query.Get("time")

	var sb strings.Builder
	sb.WriteString("Hello from " + hostName)

	if user != "" {
		sb.WriteString(", User: " + user)
	}
	if time != "" {
		sb.WriteString(", Time: " + time)
	}
	sb.WriteString("\n")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(sb.String()))
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
