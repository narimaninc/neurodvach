package server

import (
	"fmt"
	"io"
	"log"
	"math/rand/v2"
	"net/http"
	"neurodvach_backend/build"
	"neurodvach_backend/utils"
)

type handlerFunction = func(http.ResponseWriter, *http.Request) error

func registerHandlers() {
	http.HandleFunc("/", newHandler(getRoot))
}

func newHandler(fn handlerFunction) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := rand.Uint64()
		log.Printf("Got %s request at '%s' by '%s' with id '%16x'", r.Method, r.RequestURI, r.RemoteAddr, id)
		if e := fn(w, r); e != nil {
			utils.FatalfIfDebug("Falied to process request with id '%16x': %s", id, e)
		}
	}
}

func getRoot(w http.ResponseWriter, r *http.Request) error {
	io.WriteString(w, fmt.Sprintf(
		utils.Htmlize(`
		
			Neurodvach
			Version: %s
			Build configuration: %s
		
		`),
		build.Version,
		build.CurrentConfig,
	))

	return nil
}
