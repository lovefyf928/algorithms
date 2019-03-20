package main

import (
	"fmt"
	"sort"
	"time"
)

func BinaryS(key int, arr []int) (int) {
	t := time.Now()
	lo := 0
	hi := len(arr) - 1
	for hi >= lo {
		mid := lo + (hi - lo) / 2
		if arr[mid] < key {
			lo = mid + 1
		} else if arr[mid] > key {
			hi = mid - 1
		} else {
			t1 := time.Now()
			fmt.Println("BinaryS is done! from", t, " to", t1)
			return mid
		}
	}
	t1 := time.Now()
	fmt.Println("BinaryS is done! from", t, " to", t1)
	return -1
}

func NmS(key int, arr []int) int {
	t := time.Now()
	for i := 0; i < len(arr); i ++ {
		if arr[i] == key {
			t1 := time.Now()
			fmt.Println("NmS is done! from", t, " to", t1)
			return i
		}
	}
	t1 := time.Now()
	fmt.Println("NmS is done! from", t, " to", t1)
	return -1
}

func main() {
	arr := []int{}
	for i := 1; i < 100000000; i ++ {
		arr = append(arr, i)
	}
	sort.Ints(arr)
	//fmt.Println(arr)
	//fmt.Println(BinaryS(83, arr))
	go BinaryS(99999999, arr)
	go NmS(99999999, arr)
	time.Sleep(time.Second * 100)
}
