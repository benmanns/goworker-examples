package main

import (
	"fmt"
	"github.com/benmanns/goworker"
)

func init() {
	goworker.Register("Echo", echoWorker)
}

func echoWorker(queue string, args ...interface{}) error {
	fmt.Println(args)
	return nil
}
