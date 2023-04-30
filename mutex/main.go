package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("begin cpu", runtime.NumCPU)
	fmt.Println("begin num", runtime.NumGoroutine(), "\n")
	var mutex sync.Mutex
	var wg sync.WaitGroup
	var incrementer int
	incrementer = 0
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			mutex.Lock()
			i1 := incrementer
			i1++
			incrementer = i1
			fmt.Println("incr =", incrementer, "go routine -", i)
			wg.Done()
			mutex.Unlock()
		}()
	}
	//	fmt.Println("incrementer value =", incrementer)
	wg.Wait()
	fmt.Println("end value =", incrementer)
	//fmt.Println("end cpu", runtime.NumCPU)
	fmt.Println("end num", runtime.NumGoroutine(), "\n")
	//fmt.Println("The End")
}
