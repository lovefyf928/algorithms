package main

import (
	"fmt"
)

var arr = []int{13,11,10, 9, 8, 7, 6, 5, 4, 3}

func Binary(n int) (bool, interface{}) { //请输入一组降序排列的数组
	key := len(arr) / 2 - 1
	if arr[key] == n {
		return true, n
	} else {
		arr1 := arr[:key]
		arr2 := arr[key + 1:]
		if arr[key] > n {
			for {
				if len(arr2) == 0 {
					return false, nil
				}
				key1 := len(arr2) / 2 - 1
				if key1 == -1 {
					key1 = 0
				}
				if arr2[key1] == n {
					return true, n
				} else if len(arr2) != 1 {
					if arr2[key1] > n {
						arr2 = arr2[key1 + 1:]
					} else {
						arr2 = arr2[:key1]
					}
				} else {
					if arr2[0] == n {
						return true, n
					} else {
						return false, nil
					}
				}
			}
		} else {
			for {
				if len(arr1) == 0 {
					return false, nil
				}
				key2 := len(arr1) / 2 - 1
				if key2 == -1 {
					key2 = 0
				}
				if arr1[key2] == n {
					return true, n
				} else if len(arr1) != 1 {
					if arr1[key2] > n {
						arr1 = arr1[key2 + 1:]
					} else {
						arr1 = arr1[:key2]
					}
				} else {
					if arr1[0] == n {
						return true, n
					} else {
						return false, nil
					}
				}
			}
		}
	}
}

func main() {
	fmt.Println(Binary(4))
}
