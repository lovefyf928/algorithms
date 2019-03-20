package main

import "fmt"

func isPrime(N int) bool {
	if N < 2 {
		return false
	}
	for i := 2; i < N; i ++ {
		if N % i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPrime(12))
}
