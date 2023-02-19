package main

import (
	"flag"
	"fmt"
	"os"
	"selendis/modules/transfer/g"
	"selendis/modules/transfer/http"
)



func main() {

	cfg := flag.String("c", "cfg.yaml", "configuration file")
	version := flag.Bool("v", false, "show version")
	flag.Parse()

	if *version {
		fmt.Printf("version %s", Version)
		os.Exit(0)
	}

	g.ParseConfig(*cfg)

	http.Start()

	select {}
}