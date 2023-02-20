package http

import (
	"encoding/json"
	"net/http"
	"selendis/model"
	"selendis/modules/transfer/net/receiver/rpc"
)

func configApiRoutes() {
	http.HandleFunc("/api/push", apiPushDatapoints)
}

func apiPushDatapoints(rw http.ResponseWriter, req *http.Request) {
	if req.ContentLength == 0 {
		http.Error(rw, "blank body", http.StatusBadRequest)
		return
	}

	decode := json.NewDecoder(req.Body)
	var metrics []*model.MetricValue
	err := decode.Decode(&metrics)
	if err != nil {
		http.Error(rw, "decode error", http.StatusBadRequest)
		return
	}

	reply := &model.TransferResponse{}
	rpc.RecvMetricValues(metrics, reply, "http")

	RenderDataJson(rw, reply)
}
