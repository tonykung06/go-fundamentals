package main

import (
	"fmt"
)

func main() {
	// firstRank := 39
	// secondRank := 614

	if firstRank, secondRank := 39, 614; firstRank < secondRank {
		fmt.Println("The first is better than the second.")
	} else if firstRank > secondRank {
		fmt.Println("The second is better than the first.")
	} else {
		fmt.Println("equal?")
	}
}
