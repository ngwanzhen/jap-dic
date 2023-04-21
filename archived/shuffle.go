package main

import (
	"fmt"
	"math/rand"
	"time"
)

func shuffle() {
	s := []int{3, 6, 1, 8}
	for i, _ := range s {
		// generate randindex between 0 and len
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		r := r1.Intn(len(s))
		// fmt.Println(r)
		// swap val at i and r
		s[i], s[r] = s[r], s[i]
	}
	fmt.Println(s)
}
