package cron

import (
	"fmt"
	"log"
	"selendis/model"
	"selendis/modules/agent/g"
	"time"
)

func ReportAgentStatus() {
	if g.Config().Heartbeat.Enabled && g.Config().Heartbeat.Addr != "" {
		go reportAgentStatus(time.Duration(g.Config().Heartbeat.Interval) * time.Second)
	}
}

func reportAgentStatus(interval time.Duration) {
	for {
		hostname, err := g.Hostname()
		if err != nil {
			hostname = fmt.Sprintf("error:%s", err.Error())
		}

		req := model.AgentReportRequest{
			Hostname:      hostname,
			IP:            g.IP(), // IP
			AgentVersion:  g.VERSION, // 版本
			PluginVersion: "sdeas", // 插件版本号，即最后一次 git commit 的 hash
		}

		var resp model.SimpleRpcResponse
		err = g.HbsClient.Call("Agent.ReportStatus", req, &resp) // rpc 调用 hbs 的 Agent.ReportStatus，获取响应
		if err != nil || resp.Code != 0 {
			log.Println("call Agent.ReportStatus fail:", err, "Request:", req, "Response:", resp)
		}

		time.Sleep(interval) // 睡一个心跳汇报间隔时间
	}
}