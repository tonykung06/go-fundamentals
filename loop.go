package main

import (
	"fmt"
	"time"
)

func main() {
	for timer := 10; timer >= 0; timer-- {
		if timer == 0 {
			fmt.Println("Boom!")
			break
		}
		fmt.Println(timer)
		time.Sleep(1 * time.Second)
	}

	for timer := 10; timer >= 0; timer-- {
		if timer%2 == 0 {
			continue
		}
		fmt.Println(timer)
		time.Sleep(1 * time.Second)
	}

	coursesInProg := []string{"OMG", "OMG2", "OMG3"}
	coursesCompleted := []string{"hi", "OMG3", "hi3"}

FoundMatch:
	for _, value := range coursesInProg {
		fmt.Println(value)
		for _, value2 := range coursesCompleted {
			if value == value2 {
				fmt.Println("crashed on", value2)
				break FoundMatch
			}
		}
	}
}
