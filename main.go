package main

import (
	"math/rand"
	"syscall"
	"time"
)

func main() {
	go func() {
		syscall.ForkExec("echo", []string{}, &syscall.ProcAttr{})
	}()
	time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
}
