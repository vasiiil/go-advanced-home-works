package main

import (
	"fmt"
	"math"
	"math/rand/v2"
	"sync"
)

func main() {
	count := 10
	randCh := make(chan int)
	sqrCh := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		createSlice(randCh, count)
		wg.Done()
	}()
	go func() {
		sqr(randCh, sqrCh, count)
		wg.Done()
	}()
	go func() {
		wg.Wait()
		close(sqrCh)
	}()

	for num := range sqrCh {
		fmt.Print(num, " ")
	}
}

func createSlice(ch chan int, size int) {
	sl := make([]int, size)
	for i := range size {
		sl[i] = rand.IntN(100)
	}
	for _, num := range sl {
		ch <- num
	}
}
func sqr(randCh, sqrCh chan int, size int) {
	for range size {
		num := <-randCh
		sqrCh <- int(math.Pow(float64(num), 2))
	}
	close(randCh)
}
