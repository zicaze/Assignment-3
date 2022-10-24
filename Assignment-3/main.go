package main

import (
	"log"
	"microservice/handler"
	d "microservice/service"
	"net/http"
	"time"
)

func main() {

	go auto()
	http.HandleFunc("/", handler.Handler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal()
	}
}

func auto() {
	dtalert := d.GetService()
	for range time.Tick(15 * time.Second) {
		d.Updatejs(dtalert)
	}
}
