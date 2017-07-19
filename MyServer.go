package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

var mux map[string]func(http.ResponseWriter, *http.Request)

func MyServer() {
	server := http.Server{
		Addr:              ":8002",
		Handler:           &myServerHandler{},
		ReadHeaderTimeout: 5 * time.Second,
	}
	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/helloHttp"] = helloMyHttp
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

type myServerHandler struct{}

func (*myServerHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	if fun, ok := mux[request.URL.String()]; ok {
		fun(response, request)
		return
	}
	io.WriteString(response, "URL:"+request.URL.String())
}

func helloMyHttp(responseWriter http.ResponseWriter, request *http.Request) {
	io.WriteString(responseWriter, "hello My go web ! 		-wisn")
}
