package main

import (
	"fmt"
	"strconv"
)

var userInfo = map[string]int {
	"foo": 12,
	"bar": 28,
	"hello": 100,
}


func main() {
	fmt.Println(strconv.Itoa(userInfo["foo"]))
}
