package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

type Config struct{}

func main() {
	app := Config{}

	log.Printf("Service broker running on port %s\n", webPort)

	// define htpp server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// start server

	err := srv.ListenAndServe()

	if err != nil {
		log.Panic(err)
	}
}
