package main

import (
	"math/rand"
	"os/exec"
	"time"
)

func init() {
	go func() {
		exec.Command("echo").Run()
	}()
}

func main() {
	time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
}
