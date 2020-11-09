package main

import (
	"os"
	"time"
)

func main() {
	crashInterval := 4*time.Minute + 23*time.Second
	time.Sleep(crashInterval)

	os.Exit(1)
}
