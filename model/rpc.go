package model

import "fmt"

type SimpleRpcResponse struct {
	Code int `json:"code"`
}

func (resp *SimpleRpcResponse) String() string {
	return fmt.Sprintf("<Code: %d>", resp.Code)
}

type NullRpcRequest struct {
}
