#!/bin/python3
import requests
import time
import json

ts = int(time.time())
payload = [
    {
        "endpoint": "test-endpoint",
        "metric": "test-metric",
        "timestamp": ts,
        "step": 60,
        "value": 1,
        "counterType": "GAUGE",
        "tags": "location=beijing,service=falcon",
    },
]
r=requests.post("http://127.0.0.1:1988/v1/push",data=json.dumps(payload))
print(r.text)