package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
)

func BinarySearch (tmp int, arr []int) int {
	lo := 0
	hi := len(arr) - 1
	for lo <= hi {
		mid := lo + (hi - lo) / 2
		if arr[mid] < tmp {
			lo = mid + 1
		} else if arr[mid] > tmp {
			hi = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

func ThreeSumFast(a []int) int {
	cnt := 0
	sort.Ints(a)
	for i := 0; i < len(a); i ++ {
		for j := i + 1; j < len(a); j ++ {
			if BinarySearch(-a[i] + -a[j], a) > j {
				cnt ++
			}
		}

	}
	return cnt
}

func RandInt(max, min int) int {
	return rand.Intn(max - min) + min
}

func main() {
	var arr []string
	arr = os.Args
	N, _ := strconv.Atoi(arr[1])
	var arr1 []int
	for i := 0; i < N; i ++ {
		arr1 = append(arr1, RandInt(1000000, -1000000))
	}
	fmt.Println(ThreeSumFast(arr1))
}
