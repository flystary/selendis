#!/bin/bash
e="test.endpoint.1"
m="test.metric.1"
t="t0=tag0,t1=tag1,t2=tag2"
ts=$(date +%s)

curl -s -X POST -d "[{\"metric\":\"$m\", \"endpoint\":\"$e\", \"timestamp\":$ts,\"step\":60, \"value\":9, \"counterType\":\"GAUGE\",\"tags\":\"$t\"}]" "127.0.0.1:6060/api/push" | python -m json.tool