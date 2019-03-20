package main

import "os"

func IsEmpty() (bool, []interface{}) {
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
