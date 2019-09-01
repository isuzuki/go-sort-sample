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
	var len = len(ns)

	for i := 0; i < len; i++ {
		sorted := false
		for j := 0; j < len-1; j++ {
			if ns[j] > ns[j+1] {
				ns[j], ns[j+1] = ns[j+1], ns[j]
				sorted = true
			}
		}

		if !sorted {
			break
		}
	}

	return ns
}
