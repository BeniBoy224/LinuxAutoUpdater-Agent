package main

import (
	"agent/internal/logger"
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Hello, Go 🚀")
	fmt.Println("Running on:", runtime.GOOS, runtime.GOARCH)
	logger.Log("This is a log message from the logger package.")
}
