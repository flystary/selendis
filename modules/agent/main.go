package main

import (
	"flag"
	"fmt"
	"os"
	"selendis/modules/agent/cron"
	"selendis/modules/agent/funcs"
	"selendis/modules/agent/g"
	"selendis/modules/agent/http"
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

	// log level
	if g.Config().Debug {
		g.InitLog("debug")
	} else {
		g.InitLog("info")
	}

	// init
	g.InitRootDir()
	g.InitLocalIp()
	g.InitRpcClients()

	funcs.BuildMappers()

	go cron.InitDataHistory()
	cron.ReportAgentStatus()
	cron.SyncTrustableTps()

	cron.Collect()

	go http.Start()
	select {}
}
