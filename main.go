package main

import (
	"math/rand"
	"os/exec"
	"time"
)

func main() {
	go func() {
		exec.Command("echo").Run()
	}()
	time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
}
