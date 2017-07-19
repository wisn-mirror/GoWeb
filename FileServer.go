package main

import (
	"log"
	"net/http"
	"os"
)

func FileServer() {
	mux := http.NewServeMux()
	wd, err1 := os.Getwd()
	if err1 != nil {
		log.Fatal(err1)
	}
	mux.Handle("/file/", http.StripPrefix("/file/", http.FileServer(http.Dir(wd))))
	err := http.ListenAndServe(":8001", mux)
	if err != nil {
		log.Fatal(err)
	}
}
