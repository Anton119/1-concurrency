package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func square(n int) int {
	return n * n
}

// for pr

func main() {
	chIn := make(chan int, 10)
	chOut := make(chan int, 10)

	var wg sync.WaitGroup

	// создает 10 случайных чисел
	wg.Add(1)
	go func() {
		defer wg.Done()
		arr := [10]int{}
		for i := 0; i < len(arr); i++ {
			n := rand.Intn(100)
			arr[i] = n
			chIn <- arr[i]
		}
		close(chIn)

	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range chIn {
			chOut <- square(val)
		}
		close(chOut)

	}()

	wg.Wait()

	for val := range chOut {
		fmt.Print(val, " ")
	}

}
