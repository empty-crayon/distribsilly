package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/empty-crayon/distribsilly/config"
	"github.com/empty-crayon/distribsilly/db"
	"github.com/empty-crayon/distribsilly/web"
)

var (
	dbLocation = flag.String("db-location", "", "The path to the bolt db database")
	httpAddr   = flag.String("http-addr", ":8080", "HTTP Port")
	configFile = flag.String("config-file", "sharding.toml", "Config file for static sharding")
	shard      = flag.String("shard", "", "The name of the shard to store the data in")
)

func flagParse() {
	flag.Parse()

	if *dbLocation == "" {
		log.Fatalf("Must provide db-location")
	}

	if *shard == "" {
		log.Fatalf("Must provide shard")
	}
}

func main() {

	flagParse()
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.

	c, err := config.ParseFile(*configFile)

	if err != nil {
		log.Fatalf("Error parsing config %q: %v", *configFile, err)
	}

	shards, err := config.ParseShards(c.Shards, *shard)

	if err != nil {
		log.Fatalf("Error parsing shards config: %v", err)
	}

	log.Printf("Shard count is %d, current shard: %d", shards.Count, shards.CurIdx)

	db, close, err := db.NewDatabase(*dbLocation)
	if err != nil {
		log.Fatalf("Oh no!")
	}

	defer close()

	srv := web.NewServer(db, shards)

	mux := http.NewServeMux()
	mux.HandleFunc("/get", srv.GetHandler)
	mux.HandleFunc("/set", srv.SetHandler)

	err = http.ListenAndServe(*httpAddr, mux)
	if err != nil {
		log.Fatal(err)
	}
}
