package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	switch "docker" {
	case "docker":
		fmt.Println("docker")
		fallthrough
	case "linux":
		fmt.Println("linux")
	case "windows":
		fmt.Println("windows")
	default:
		fmt.Println("no match")
	}

	switch tmpNum := random(); tmpNum {
	case 0, 2, 4, 6, 8:
		fmt.Println("even number")
	case 1, 3, 5, 7, 9:
		fmt.Println("even number")
	}
}

func random() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(10)
}
