package main

import (
	"fmt"
	"github.com/MikeBazh/dog"
	"github.com/MikeBazh/puppy"
)

func main() {
	s1 := puppy.Bark()
	fmt.Println(s1)
	s2 := puppy.BigBark()
	fmt.Println(s2)
	fmt.Println("1 human year is", puppy.Years(1), "dog years")
	fmt.Println(dog.Years(1))
}
