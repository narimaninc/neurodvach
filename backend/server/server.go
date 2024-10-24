package server

import (
	"net/http"
	"sync"
)

var Server *http.Server = &http.Server{}
var serverError error

var serverWait *sync.WaitGroup = &sync.WaitGroup{}

func InitializeServer(address string) {
	Server.Addr = address

	registerHandlers()
}

func RunServer() {
	serverWait.Add(1)
	go func() {
		defer serverWait.Done()
		serverError = Server.ListenAndServe()
	}()
}

func WaitServer() error {
	serverWait.Wait()
	return serverError
}
