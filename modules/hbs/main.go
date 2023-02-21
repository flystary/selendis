package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"selendis/modules/hbs/g"
	"selendis/modules/hbs/rpc"
	"syscall"
)

var (
	Version    = "<UNDEFINED>"
	GitCommit  = "<UNDEFINED>"
	BinaryName = "<UNDEFINED>"
)

func main() {

	g.Version = Version

	cfg := flag.String("c", "cfg.yaml", "configuretion file")
	version := flag.Bool("v", false, "Display version")
	flag.Parse()

	if *version {
		fmt.Printf("version %s", Version)
		os.Exit(0)
	}

	g.ParseConfig(*cfg)

	go rpc.Start()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		fmt.Println()
		// db.DB.Close()
		os.Exit(0)
	}()

	select {}

}
