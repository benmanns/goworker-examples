package main

import (
	"fmt"
	"github.com/benmanns/goworker"
	"time"
)

func init() {
	goworker.Register("Sleep", sleepWorker)
}

func sleepWorker(queue string, args ...interface{}) error {
	seconds, ok := args[0].(float64)
	if !ok {
		return fmt.Errorf("Invalid parameters %v to sleep worker. Expected int, was %T.", args, args[0])
	}
	time.Sleep(time.Duration(int(seconds * float64(time.Second))))
	return nil
}
