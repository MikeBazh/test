package main

import "fmt"

func main() {
	ii := []int{1, 2, 3, 4, 5, 6}
	foo(ii)

}
func foo(ii []int) {
	sum := 0
	for _, v := range ii {
		sum = +v
	}
	fmt.Println(sum)
}
