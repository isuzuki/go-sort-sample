package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var ns = make([]int, n)
	for i := 0; i < len(ns); i++ {
		fmt.Scan(&ns[i])
	}

	fmt.Println(sort(ns))
}

func sort(ns []int) []int {
	for i := 1; i < len(ns); i++ {
		for j := 0; j < i+1; j++ {
			if ns[i] < ns[j] {
				ns[j], ns[i] = ns[i], ns[j]
			}
		}
	}

	return ns
}
