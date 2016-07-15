package main

import (
	"fmt"
)

func main() {
	leagueTtitles2 := make(map[string]int, 10) //for big map, give a size for performance
	fmt.Println(leagueTtitles2)
	leagueTtitles := make(map[string]int)
	leagueTtitles["Sunderland"] = 6
	leagueTtitles["Newcastle"] = 4

	recentHead2Head := map[string]int{
		"Sunderland": 5,
		"Newcastle":  0,
	}

	fmt.Printf("\nLeague titles: %v\nRecent head to heads: %v\n", leagueTtitles, recentHead2Head)

	testMap := map[string]int{"A": 1, "B": 2, "C": 3, "D": 4, "E": 5}
	testMap["C"] = 10
	testMap["F"] = 102
	delete(testMap, "A")
	fmt.Println(testMap)
	for key, value := range testMap {
		fmt.Printf("\nKey is: %v Value is: %v", key, value)
	}
}
