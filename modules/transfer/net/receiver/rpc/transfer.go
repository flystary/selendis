package rpc

import (
	"fmt"
	"log"
	"selendis/model"
	"time"
)

type Transfer int

type TransferResp struct {
	Msg        string
	Total      int
	ErrInvalid int
	Latency    int64
}

func (resp *TransferResp) String() string {
	s := fmt.Sprintf("TransferResp total=%d, err_invalid=%d, latency=%dms",
		resp.Total, resp.ErrInvalid, resp.Latency)
	if resp.Msg != "" {
		s = fmt.Sprintf("%s, msg=%s", s, resp.Msg)
	}
	return s
}

func (t *Transfer) Ping(args []*model.MetricValue, reply *model.TransferResponse) error {
	return nil
}

func (t *Transfer) Update(args []*model.MetricValue, reply *model.TransferResponse) error {
	return RecvMetricValues(args, reply, "rpc")
}

func RecvMetricValues(args []*model.MetricValue, reply *model.TransferResponse, from string) error {
	start := time.Now()
	reply.Invalid = 0

	// items := []*model.MetaData{}

	log.Println(args)
	reply.Message = "ok"
	reply.Total = len(args)
	reply.Latency = (time.Now().UnixNano() - start.UnixNano()) / 1000000

	return nil
}
