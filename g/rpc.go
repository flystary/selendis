package g

import (
	"log"
	"math"
	"net/rpc"
	"sync"
	"time"

	"github.com/toolkits/net"
)

type SingleConnRpcClient struct {
	sync.Mutex
	rpcClient 		*rpc.Client
	RpcServer 		string
	Timeout   		time.Duration
}

func (src *SingleConnRpcClient) close() {
	if src.rpcClient != nil {
		src.rpcClient.Close()
		src.rpcClient = nil
	}
}

func (src *SingleConnRpcClient) connServer() error {
	if src.rpcClient != nil {
		return nil
	}

	var err error
	var retry int = 1

	for {
		if src.rpcClient != nil {
			return nil
		}
		src.rpcClient, err = net.JsonRpcClient("tcp", src.RpcServer, src.Timeout)
		if err != nil {
			log.Printf("dial %s fail: %v", src.RpcServer, err)
			if retry > 3 {
				return err
			}
		}
		time.Sleep(time.Duration(math.Pow(2.0, float64(retry))) * time.Second)
		retry++
		continue
	}
	return err
}
