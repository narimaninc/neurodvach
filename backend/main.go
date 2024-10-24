package main

import (
	"log"
	"net/http"
	"neurodvach_backend/build"
	"neurodvach_backend/server"
	"neurodvach_backend/utils"
)

func main() {
	server.InitializeServer(":1337")
	server.RunServer()

	if build.Debug {
		if e := utils.OpenBrowser("http://localhost:1337"); e != nil {
			log.Fatal(e)
		}
	}

	if e := server.WaitServer(); e == http.ErrServerClosed {
		log.Print("Server stopped gracefully")
	} else {
		log.Printf("Server stopped with error: %s", e)
	}
}
