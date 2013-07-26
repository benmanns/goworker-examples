package main

import (
	"github.com/benmanns/goworker"
)

func init() {
	goworker.Register("Nil", nilWorker)
}

func nilWorker(queue string, args ...interface{}) error {
	return nil
}
