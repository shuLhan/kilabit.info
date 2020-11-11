package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	version := "v0.1.0"
	x := 0
	for {
		fmt.Printf("stdout: %s: %d\n", version, x)
		log.Printf("stderr: %s: %d\n", version, x)
		x++
		time.Sleep(3 * time.Second)
	}
}
