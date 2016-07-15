package main

import (
	"fmt"
)

func main() {
	type courseMeta struct {
		Author string
		Level  string
		Rating float64
	}

	var DockerDeepDive1 courseMeta
	fmt.Println(DockerDeepDive1)
	DockerDeepDive2 := new(courseMeta)
	fmt.Println(DockerDeepDive2)
	DockerDeepDive3 := courseMeta{
		Author: "Tony",
		Level:  "Intermediate",
		Rating: 5,
	}
	DockerDeepDive3.Rating = 3
	fmt.Println(DockerDeepDive3)
	fmt.Printf("%T, %v", DockerDeepDive3.Rating, DockerDeepDive3.Rating)
}
