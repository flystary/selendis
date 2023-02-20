package proc

import (
	"log"

	"github.com/toolkits/proc"
)

var (
	RecvDataTrace = proc.NewDataTrace("RecvDataTrace", 3)
)

var (
	RecvDataFiler = proc.NewDataFilter("RecvDataFilter", 5)
)

// 统计指标的整体数据
var (
	// 计数统计,正确计数,错误计数, ...
	RecvCnt       = proc.NewSCounterQps("RecvCnt")
	RpcRecvCnt    = proc.NewSCounterQps("RpcRecvCnt")
	HttpRecvCnt   = proc.NewSCounterQps("HttpRecvCnt")
	SocketRecvCnt = proc.NewSCounterQps("SocketRecvCnt")

	SendToJudgeCnt    = proc.NewSCounterQps("SendToJudgeCnt")
	SendToTsdbCnt     = proc.NewSCounterQps("SendToTsdbCnt")
	SendToGraphCnt    = proc.NewSCounterQps("SendToGraphCnt")
	SendToTransferCnt = proc.NewSCounterQps("SendToTransferCnt")
	SendToInfluxdbCnt = proc.NewSCounterQps("SendToInfluxdbCnt")
	SendToP8sRelayCnt = proc.NewSCounterQps("SendToP8sRelayCnt")

	SendToJudgeDropCnt    = proc.NewSCounterQps("SendToJudgeDropCnt")
	SendToTsdbDropCnt     = proc.NewSCounterQps("SendToTsdbDropCnt")
	SendToGraphDropCnt    = proc.NewSCounterQps("SendToGraphDropCnt")
	SendToTransferDropCnt = proc.NewSCounterQps("SendToTransferDropCnt")
	SendToInfluxdbDropCnt = proc.NewSCounterQps("SendToInfluxdbDropCnt")
	SendToP8sRelayDropCnt = proc.NewSCounterQps("SendToP8sRelayDropCnt")

	SendToJudgeFailCnt    = proc.NewSCounterQps("SendToJudgeFailCnt")
	SendToTsdbFailCnt     = proc.NewSCounterQps("SendToTsdbFailCnt")
	SendToGraphFailCnt    = proc.NewSCounterQps("SendToGraphFailCnt")
	SendToTransferFailCnt = proc.NewSCounterQps("SendToTransferFailCnt")
	SendToInfluxdbFailCnt = proc.NewSCounterQps("SendToInfluxdbFailCnt")
	SendToP8sRelayFailCnt = proc.NewSCounterQps("SendToP8sRelayFailCnt")

	// 发送缓存大小
	JudgeQueuesCnt    = proc.NewSCounterBase("JudgeSendCacheCnt")
	TsdbQueuesCnt     = proc.NewSCounterBase("TsdbSendCacheCnt")
	GraphQueuesCnt    = proc.NewSCounterBase("GraphSendCacheCnt")
	TransferQueuesCnt = proc.NewSCounterBase("TransferSendCacheCnt")
	P8sRelayQueuesCnt = proc.NewSCounterBase("P8sRelaySendCacheCnt")

	// http请求次数
	HistoryRequestCnt = proc.NewSCounterQps("HistoryRequestCnt")
	InfoRequestCnt    = proc.NewSCounterQps("InfoRequestCnt")
	LastRequestCnt    = proc.NewSCounterQps("LastRequestCnt")
	LastRawRequestCnt = proc.NewSCounterQps("LastRawRequestCnt")

	// http回执的监控数据条数
	HistoryResponseCounterCnt = proc.NewSCounterQps("HistoryResponseCounterCnt")
	HistoryResponseItemCnt    = proc.NewSCounterQps("HistoryResponseItemCnt")
	LastRequestItemCnt        = proc.NewSCounterQps("LastRequestItemCnt")
	LastRawRequestItemCnt     = proc.NewSCounterQps("LastRawRequestItemCnt")
)

func Start() {
	log.Println("proc.Start, ok")
}

func GetAll() []interface{} {
	ret := make([]interface{}, 0)

	// recv cnt
	ret = append(ret, RecvCnt.Get())
	ret = append(ret, RpcRecvCnt.Get())
	ret = append(ret, HttpRecvCnt.Get())
	ret = append(ret, SocketRecvCnt.Get())

	// send cnt
	ret = append(ret, SendToJudgeCnt.Get())
	ret = append(ret, SendToTsdbCnt.Get())
	ret = append(ret, SendToInfluxdbCnt.Get())
	ret = append(ret, SendToGraphCnt.Get())
	ret = append(ret, SendToTransferCnt.Get())
	ret = append(ret, SendToP8sRelayCnt.Get())

	// drop cnt
	ret = append(ret, SendToJudgeDropCnt.Get())
	ret = append(ret, SendToTsdbDropCnt.Get())
	ret = append(ret, SendToGraphDropCnt.Get())
	ret = append(ret, SendToTransferDropCnt.Get())
	ret = append(ret, SendToInfluxdbDropCnt.Get())
	ret = append(ret, SendToP8sRelayDropCnt.Get())

	// send fail cnt
	ret = append(ret, SendToJudgeFailCnt.Get())
	ret = append(ret, SendToTsdbFailCnt.Get())
	ret = append(ret, SendToGraphFailCnt.Get())
	ret = append(ret, SendToTransferFailCnt.Get())
	ret = append(ret, SendToInfluxdbFailCnt.Get())
	ret = append(ret, SendToP8sRelayFailCnt.Get())

	// cache cnt
	ret = append(ret, JudgeQueuesCnt.Get())
	ret = append(ret, TsdbQueuesCnt.Get())
	ret = append(ret, GraphQueuesCnt.Get())
	ret = append(ret, TransferQueuesCnt.Get())
	ret = append(ret, P8sRelayQueuesCnt.Get())

	// http request
	ret = append(ret, HistoryRequestCnt.Get())
	ret = append(ret, InfoRequestCnt.Get())
	ret = append(ret, LastRequestCnt.Get())
	ret = append(ret, LastRawRequestCnt.Get())

	// http response
	ret = append(ret, HistoryResponseCounterCnt.Get())
	ret = append(ret, HistoryResponseItemCnt.Get())
	ret = append(ret, LastRequestItemCnt.Get())
	ret = append(ret, LastRawRequestItemCnt.Get())

	return ret
}