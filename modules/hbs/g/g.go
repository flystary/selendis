package g

import (
	"log"
	"runtime"
)

var (
	BinaryName string = "hbs"
	Version    string = "1.0.0"
	GitCommit  string = "hsbas"
)

func VersionMsg() string {
	return Version + "@" + GitCommit
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}