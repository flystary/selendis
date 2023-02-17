package main

import (
	"flag"
	"fmt"
	"os"
	"selendis/g"
)


func main() {
	//cfg 
	cfg := flag.String("c", "cfg.yaml", "configuretion file")
	version := flag.Bool("v", false, "display Version")

	flag.Parse()

	// print Version
	if *version {
		fmt.Println(g.VERSION)
		os.Exit(0)
	}

	// load cfg
	g.ParseConfig(*cfg)


	select {}
}