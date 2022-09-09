package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {
	//Get x from command line
	args := os.Args[1:]
	var n int
	n, _ = strconv.Atoi(args[0])

	//Generate x random numbers in slice
	s1 := make([]int, n)
	for i := 0; i < n; i++ {
		s1[i] = rand.Intn(100)
	}

	//Use sort.Slice and sort.SliceStable to sort the slice
	//Output the time it takes to sort the slice
	start := time.Now()

	sort.Slice(s1, func(i, j int) bool { //https://stackoverflow.com/questions/37695209/golang-sort-slice-ascending-or-descending
		return s1[i] < s1[j]
	})

	end := time.Now()
	runtime := end.Sub(start)
	fmt.Println("Slice Sort Runtime: ", runtime)

	s2 := make([]int, n)
	for i := 0; i < n; i++ {
		s2[i] = rand.Intn(100)
	}

	start = time.Now()

	sort.SliceStable(s1, func(i, j int) bool { //https://stackoverflow.com/questions/37695209/golang-sort-slice-ascending-or-descending
		return s1[i] < s1[j]
	})

	end = time.Now()
	runtime = end.Sub(start)
	fmt.Println("Slice Stable Runtime: ", runtime)

}
