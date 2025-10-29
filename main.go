package main

import (
	"fmt"
	"math"
	"math/rand/v2"
	"sync"
)

func main() {
	count := 10
	randCh := make(chan int, count)
	sqrCh := make(chan int, count)
	result := make([]int, count)
	var wg sync.WaitGroup
	for i := range count {
		wg.Add(3)
		go func() {
			createRandInt(randCh)
			wg.Done()
		}()
		go func() {
			sqr(randCh, sqrCh)
			wg.Done()
		}()
		go func(i int) {
			num := <- sqrCh
			result[i] = num
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println(result)
}

func createRandInt(ch chan int) {
	ch <- rand.IntN(100)
}
func sqr(randCh, sqrCh chan int) {
	num := <- randCh
	sqrCh <- int(math.Pow(float64(num), 2))
}
