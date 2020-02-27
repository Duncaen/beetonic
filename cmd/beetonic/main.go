package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Duncaen/beetonic/subsonic"
	"github.com/Duncaen/beetonic/beets"

	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

func main() {
	b, err := beets.Open("/home/duncan/data/beets/library.db")
	if err != nil {
		log.Fatal(err)
	}
	defer b.Close()
	s, err := subsonic.New(subsonic.Library(b))
	if err != nil {
		log.Fatal(err)
	}
	r := mux.NewRouter()
	s.RegisterRoutes(r)
	log.Fatal(http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stderr, r)))
}
