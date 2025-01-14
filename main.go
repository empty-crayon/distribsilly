package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/empty-crayon/distribsilly/db"
	"github.com/empty-crayon/distribsilly/web"
)

var (
	dbLocation = flag.String("db-location", "", "The path to the bolt db database")
	httpAddr   = flag.String("http-addr", ":8080", "HTTP Port")
)

func flagParse() {
	flag.Parse()

	if *dbLocation == "" {
		log.Fatalf("Must provide db-location")
	}
}

func main() {

	flagParse()
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.

	db, close, err := db.NewDatabase(*dbLocation)

	if err != nil {
		log.Fatalf("Oh no!")
	}

	defer close()

	srv := web.NewServer(db)

	mux := http.NewServeMux()
	mux.HandleFunc("/get", srv.GetHandler)
	mux.HandleFunc("/set", srv.SetHandler)

	err = http.ListenAndServe(*httpAddr, mux)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hello from my world!")

}
