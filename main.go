package main

import (
	"log"

	"github.com/asdine/storm"
	"github.com/namsral/flag"
)

var (
	cfg Config
	db  *storm.DB
)

func main() {
	var (
		config string
		data   string
		dbpath string
		bind   string
	)

	flag.StringVar(&config, "config", "", "config file")
	flag.StringVar(&data, "data", "./data", "path to data")
	flag.StringVar(&dbpath, "dbpath", "notes.db", "Database path")
	flag.StringVar(&bind, "bind", "0.0.0.0:8000", "[int]:<port> to bind to")
	flag.Parse()

	var err error
	db, err = storm.Open(dbpath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// TODO: Abstract the Config and Handlers better
	cfg.data = data

	NewServer(bind, cfg).ListenAndServe()
}
