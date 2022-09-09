package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func sum(a []int, ch chan int) {
	s := 0
	for _, v := range a {
		s += v
	}
	ch <- s
}

func main() {
	args := os.Args[1:]
	var n int
	n, _ = strconv.Atoi(args[0])

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(100)
	}

	s := 0
	start := time.Now()

	for _, v := range arr {
		s += v
	}

	end := time.Now()
	runtime := end.Sub(start)

	fmt.Println("Sum: ", s)
	fmt.Println("Normal Runtime: ", runtime)

	ch := make(chan int)

	start = time.Now()

	go sum(arr[:len(arr)/2], ch)
	go sum(arr[len(arr)/2:], ch)

	s1, s2 := <-ch, <-ch

	end = time.Now()
	runtime = end.Sub(start)

	fmt.Println("Sum: ", s1+s2)
	fmt.Println("Parallel Runtime: ", runtime)
}
