package main

import (
	"fmt"
	"strconv"

	"main.go/internal/cache"
)

const CAPACITY = 4

func checkHit(c cache.LRUI, key cache.Key) {
	val, ok := c.Get(key)

	if !ok {
		fmt.Println("no value for key: ", key)
	} else {
		fmt.Println(key, " -> ", val)
	}

}

func main() {
	cap := 10

	fmt.Println("SLICE:")
	lruS, err := cache.NewLRU("slice", cap)

	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 1; i <= cap+1; i++ {
		val := "abc" + strconv.Itoa(i)
		lruS.Add(i, val)
	}

	checkHit(lruS, 0)
	checkHit(lruS, 1)
	checkHit(lruS, 9)
	checkHit(lruS, 10)
	checkHit(lruS, 11)
	checkHit(lruS, 12)

	fmt.Println("LIST:")

	lru, err := cache.NewLRU("list", cap)

	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 1; i <= cap+1; i++ {
		val := "xyz" + strconv.Itoa(i)
		lru.Add(i, val)
	}

	checkHit(lru, 0)
	checkHit(lru, 1)
	checkHit(lru, 9)
	checkHit(lru, 10)
	checkHit(lru, 11)
	checkHit(lru, 12)

}
