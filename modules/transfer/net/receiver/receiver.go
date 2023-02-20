package receiver

import "selendis/modules/transfer/net/receiver/rpc"

func Start() {
	go rpc.StartRpc()
}
