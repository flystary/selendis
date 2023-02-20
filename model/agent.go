package model

import "fmt"

type AgentReportRequest struct {
	Hostname string
	IP		 string
	AgentVersion	string
	PluginVersion	string
}

func (this *AgentReportRequest) String() string {
	return fmt.Sprintf(
		"<Hostname:%s, IP:%s, AgentVersion:%s, PluginVersion:%s>",
		this.Hostname,
		this.IP,
		this.AgentVersion,
		this.PluginVersion,
	)
}

type AgentUpdateInfo struct {
	LastUpdate    int64
	ReportRequest *AgentReportRequest
}