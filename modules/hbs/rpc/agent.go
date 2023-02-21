package rpc

import (
	"selendis/model"
	"selendis/modules/hbs/g"
	"strings"
)


func (t *Agent) ReportStatus(args *model.AgentReportRequest, reply *model.SimpleRpcResponse) error {
	if args.Hostname == "" {
		reply.Code = 1
		return nil
	}

	// cache.Agents.Put(args)

	return nil
}

func (t *Agent) TrustableIps(args *model.NullRpcRequest, ips *string) error {
	*ips = strings.Join(g.Config().Trustable, ",")
	return nil
}