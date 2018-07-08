package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	fmt.Println("Hello World.")
	fmt.Println("System Info:")
	fmt.Printf("Go Version: %s\n", runtime.Version())
	fmt.Printf("Num Goroutine: %d\n", runtime.NumGoroutine())
	fmt.Printf("Num CPU: %d\n", runtime.NumCPU())
	fmt.Printf("Num Cgo Call: %d\n", runtime.NumCgoCall())
	hostname, _ := os.Hostname()
	fmt.Printf("Hostname: %s\n", hostname)
}
