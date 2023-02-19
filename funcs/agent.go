package funcs

import "selendis/model"


func AgentMetrics() (mms []*model.MetricValue) {
	return []*model.MetricValue{GaugeValue("agent.alive", 1)}
}