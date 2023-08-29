package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//xs := []string{"in", "my", "younger"}
	//	fmt.Println(xs[2])
	wordfreq := map[string]int{}
	data1 := []string{}
	readFile, err := os.Open("text.txt")
	if err != nil {
		fmt.Println(err)
	}
	s := bufio.NewScanner(readFile)
	s.Split(bufio.ScanWords)
	for s.Scan() {
		data := s.Text()
		fmt.Printf("%v\t", data)
		fmt.Printf("\n")
		data = strings.Replace(data, `"`, "", -1)
		data = strings.Replace(data, `,`, "", -1)
		data1 = append(data1, data)
		wordfreq[data]++
	}
	fmt.Println("__________")
	//fmt.Println(data1)
	for k, v := range wordfreq {
		fmt.Println(k, ":", v)
	}
}
