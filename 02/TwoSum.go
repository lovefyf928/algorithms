package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func RandInt(min, max int) int {
	return rand.Intn(max - min) + min
}

func count(a []int) int {
	cnt := 0
	for i := 0; i < len(a); i ++ {
		for j := i + 1; j < len(a); j ++ {
			if a[i] + a[j] == 0 {
				cnt ++
			}
		}
	}
	return cnt
}

func main() {
	var arr []int
	var ars []string
	ars = os.Args
	N, _ := strconv.Atoi(ars[1])
	for i := 0; i < N; i ++ {
		arr = append(arr, RandInt(-1000000, 1000000))
	}
	fmt.Println(count(arr))
}
