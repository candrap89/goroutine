package main

import (
	"fmt"
	"sync"
)

func prime(n int, output chan []int, wg *sync.WaitGroup) {
	defer wg.Done()
	var prime []int
	for i := 0; i < n; i++ {
		if isPrima(i) {
			prime = append(prime, i)
		}
	}
	output <- prime
}

func isPrima(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func fibonaci(n int, ch chan []int, wg *sync.WaitGroup) {
	defer wg.Done()
	if n <= 0 {
		ch <- []int{}
	}

	fib := make([]int, n)
	fib[0] = 0
	if n > 1 {
		fib[1] = 1
	}
	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	ch <- fib

}

func main() {

	// Channel untuk komunikasi antar goroutine
	primeCh := make(chan []int)
	fibCh := make(chan []int)

	// contoh concurency
	var wg sync.WaitGroup
	wg.Add(2)

	// Menjalankan fungsi dalam goroutine
	go prime(10, primeCh, &wg)
	go fibonaci(10, fibCh, &wg)

	// Menunggu goroutine selesai
	go func() {
		wg.Wait()
		close(primeCh)
		close(fibCh)
	}()

	// Menerima hasil dari channel
	primes := <-primeCh
	fibonacci := <-fibCh

	// Menampilkan hasil
	fmt.Println("Bilangan Prima:", primes)
	fmt.Println("Bilangan Fibonacci:", fibonacci)

}

// input list of number
// output nya list
