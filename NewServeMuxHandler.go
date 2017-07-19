package main

import (
	"io"
	"log"
	"net/http"
)

func NewSerMuxHandler() {
	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{})
	mux.HandleFunc("/helloHttp", helloHttp)
	err := http.ListenAndServe(":8001", mux)
	if err != nil {
		log.Fatal(err)
	}
}

type myHandler struct{}

func (*myHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	io.WriteString(responseWriter, " go Web "+request.URL.String()+" "+request.Referer())
}

func helloHttp(responseWriter http.ResponseWriter, request *http.Request) {
	io.WriteString(responseWriter, "hello go web ! 		-wisn")
}
