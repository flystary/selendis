package g

import (
	"bytes"
	"log"
	"net"
	"os"
	"selendis/model"
	"strings"
	"time"

	"sync"

	"github.com/toolkits/slice"
)

var Root string

func InitRootDir() {
	var err error
	Root, err = os.Getwd()
	if err != nil {
		log.Fatalln("getwd fail:", err)
	}
}

var LocalIp string

func InitLocalIp() {
	if Config().Heartbeat.Enabled {
		conn, err := net.DialTimeout("tcp", Config().Heartbeat.Addr, time.Second*10)
		if err != nil {
			log.Println("get local addr failed !")
		} else {
			LocalIp = strings.Split(conn.LocalAddr().String(), ":")[0]
			conn.Close()
		}
	} else {
		log.Println("hearbeat is not enabled, can't get Localip")
	}
}

var (
	HbsClient *SingleConnRpcClient
)

func InitRpcClients() {
	if Config().Heartbeat.Enabled {
		HbsClient = &SingleConnRpcClient{
			RpcServer: Config().Heartbeat.Addr,
			Timeout:   time.Duration(config.Heartbeat.Timeout) * time.Millisecond,
		}
	}
}

func SendToTransfer(metrics []*model.MetricValue) {
	if len(metrics) == 0 {
		return
	}
	dftags := Config().DefaultTags

	if len(dftags) > 0 {
		var buf bytes.Buffer
		default_tags_list := []string{}

		for k, v := range dftags {
			buf.Reset()
			buf.WriteString(k)
			buf.WriteString("=")
			buf.WriteString(v)

			default_tags_list = append(default_tags_list, buf.String())
		}
		default_tags := strings.Join(default_tags_list, ",")
		for i, x := range metrics {
			buf.Reset()
			if x.Tags == "" {
				metrics[i].Tags = default_tags
			} else {
				buf.WriteString(metrics[i].Tags)
				buf.WriteString(",")
				buf.WriteString(default_tags)
				metrics[i].Tags = buf.String()
			}
			log.Print("xxwTest:")
			log.Println(metrics[i])
		}
	}

	debug := Config().Debug
	if debug {
		log.Printf("=> <Total=%d> %v\n", len(metrics), metrics[0]) // debug 模式下，记录每次上报 metric 的个数并记录第一个 metric
	}

	var resp model.TransferResponse
	SendMetrics(metrics, &resp) // 发送给 transfer

	if debug {
		log.Println("<=", &resp)
	}
}

var (
	ips     []string
	ipsLock = new(sync.Mutex)
)

func TrustableIps() []string {
	ipsLock.Lock()
	defer ipsLock.Unlock()
	return ips
}

func SetTrustableIps(ipStr string) {
	arr := strings.Split(ipStr, ",")
	ipsLock.Lock()
	defer ipsLock.Unlock()
	ips = arr
}

func IsTrustable(remoteAddr string) bool {
	ip := remoteAddr
	idx := strings.LastIndex(remoteAddr, ":")
	if idx > 0 {
		ip = remoteAddr[0:idx]
	}

	if ip == "127.0.0.1" {
		return true
	}

	return slice.ContainsString(TrustableIps(), ip)
}
