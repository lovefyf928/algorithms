package main

import "fmt"

func main() {
	var arr = [4]float32{}
	for i := 0; i < 4; i ++ {
		arr[i] = 0.0
	}
	fmt.Println(arr)
}
