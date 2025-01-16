package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/BurntSushi/toml"
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

	var c config.Config
	if _, err := toml.DecodeFile(*configFile, &c); err != nil {
		log.Fatalf("Error decoding toml file, toml.decodefile(%q): %v", *configFile, err)
	}
	

	shards, err := config.ParseShards(c.Shards, *shard)


	// what is the shard count and what is the shard idx
	var shardCount int = len(c.Shards)
	var shardIdx = -1
	var addrs = make(map[int]string)

	for _, s := range c.Shards {
		// populate the address map for the shards
		addrs[s.Idx] = s.Address

		if s.Name == *shard {
			shardIdx = s.Idx
		}
	}

	if shardIdx < 0 {
		log.Fatalf("Shard %q not found!", *shard)
	}



	log.Printf("Shard count is %d, currently writing to shard: %d", shardCount, shardIdx)
	db, close, err := db.NewDatabase(*dbLocation)

	if err != nil {
		log.Fatalf("Oh no!")
	}

	defer close()

	srv := web.NewServer(db, shardIdx, shardCount, addrs)

	mux := http.NewServeMux()
	mux.HandleFunc("/get", srv.GetHandler)
	mux.HandleFunc("/set", srv.SetHandler)

	err = http.ListenAndServe(*httpAddr, mux)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hello from my world!")

}
