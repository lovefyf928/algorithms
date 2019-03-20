package main

import (
	"fmt"
)

func rank(key int, a []int) int {
	return rank1(key, a, 0, len(a) - 1)
}

func rank1(key int, a []int, lo int, hi int) int {
	if lo > hi {
		return -1
	}
	mid := lo + (hi - lo) / 2
	if key > a[mid] {
		return rank1(key, a, lo, mid - 1)
	} else if key < a[mid] {
		return rank1(key, a, mid + 1, hi)
	} else {
		return mid
	}
}

func main() {
	set := []int{11, 10, 9, 8, 7, 6, 5 ,4, 3}
	//sort.Ints(set)
	fmt.Println(rank(4, set))
	fmt.Printf("%14d", rank(4, set))
}
