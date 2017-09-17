package conprng

import (
	"fmt"
)

func generate(ch chan int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

func filter(in, out chan int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

func generatePrimes(num int) []int {
	inChan := make(chan int)
	primes := make([]int, 0, num)
	go generate(inChan)
	for len(primes) < num {
		prime := <-inChan
		primes = append(primes, prime)
		outChan := make(chan int)
		go filter(inChan, outChan, prime)
		inChan = outChan
	}
	return primes
}
