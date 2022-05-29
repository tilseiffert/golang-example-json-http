package main

import (
	"fmt"
	"time"
	"github.com/tilseiffert/golang-example-json-http/internal/server"
)

func debug(s string) {
    currentTime := time.Now().String()

	fmt.Printf("%s DEBUG: %s\n", currentTime, s)
}



func main() {
	debug("Hello World!")

	srv := server.NewHTTPServer(":8080")

/*
	debug("Setting up demo data")
	demo := server.Activity{Time: time.Now(), Description: "Example Description"}	
	fmt.Printf("demo id %d inserted \n", server.LastActivities.Insert(demo))
	fmt.Printf("demo id %d inserted \n", server.LastActivities.Insert(demo))
	fmt.Printf("demo id %d inserted \n", server.LastActivities.Insert(demo))
*/
	
	debug("Server startet, listening on port 8080 :D")
	srv.ListenAndServe()
}
