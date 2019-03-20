package main

import (
	"fmt"
	"os"
)

func isEmpty() (bool, []interface{}) {
	list := os.Args
	if len(list) > 0 {
		arr := []interface{}{}
		for i := 0; i < len(list); i ++ {
			arr = append(arr, list[i])
		}
		return true, arr
	}
	return false, nil
}


func main() {
	ok, arr := isEmpty()
	if ok {
		for i := 1; i < len(arr); i ++ {
			if i == len(arr) - 1 {
				fmt.Println("equal")
				break
			}
			if arr[i] != arr[i + 1] {
				fmt.Println("not equal")
				break
			}
		}
	}
}
