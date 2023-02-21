package rpc

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"selendis/modules/hbs/g"
	"time"
)

type Hbs int
type Agent int


func Start() {
	addr := g.Config().Listen

	server := rpc.NewServer()
	server.Register(new(Agent))
	server.Register(new(Hbs))

	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("listen error:", err)
	} else {
		log.Println("listening", addr)
	}

	for {
		for {
			conn, err := listen.Accept()
			if err != nil {
				log.Println("listener accept fail:", err)
				time.Sleep(time.Duration(100) * time.Millisecond)
				continue
			}
			go server.ServeCodec(jsonrpc.NewServerCodec(conn))
		}
	}
}