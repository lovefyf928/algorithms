package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

type Counter interface {
	increment()
	tally() int
	toString() string
}

func (q *q) increment(){
	q.count ++
}

type q struct {
	id string
	count int
	Counter
}


func main() {
	var heads q
	var tails q
	heads.id = "heads"
	tails.id = "tails"
	var list []string
	list = os.Args
	T, err := strconv.Atoi(list[1])
	if err == nil {
		for i := 0; i < T; i ++ {
			r := rand.Intn(2)
			if r == 0 {
				heads.increment()
			} else {
				tails.increment()
			}
		}
		fmt.Println(heads.id, heads.count)
		fmt.Println(tails.id, tails.count)
	}
}
