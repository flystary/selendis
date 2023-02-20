package main

import (
	"flag"
	"fmt"
	"os"
	"selendis/modules/transfer/g"
	"selendis/modules/transfer/http"
	"selendis/modules/transfer/net/receiver"
	"selendis/modules/transfer/tools/proc"
)

func main() {

	g.Version = Version

	cfg := flag.String("c", "cfg.yaml", "configuration file")
	version := flag.Bool("v", false, "show version")
	flag.Parse()

	if *version {
		fmt.Printf("version %s", Version)
		os.Exit(0)
	}

	g.ParseConfig(*cfg)
	proc.Start()

	receiver.Start()

	http.Start()

	select {}
}
