package main

import (
	"anog/handlers"
	"github.com/dropbox/godropbox/errors"
	"github.com/valyala/fasthttp"
	"log"
)

func main() {
	router := handlers.NewRouter(handlers.NewAnagram(), handlers.NewLoadPrepare())
	server := &fasthttp.Server{
		Handler:        router.Route,
	}
	err := server.ListenAndServe(":8080")
	if err != nil {
		log.Fatal(errors.New(err.Error()))
	}
}

