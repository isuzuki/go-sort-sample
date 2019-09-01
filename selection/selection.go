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
	len := len(ns)
	for i := 0; i < len-1; i++ {
		n := ns[i]
		m := i
		for j := i + 1; j < len; j++ {
			if n > ns[j] {
				m = j
				n = ns[j]
			}
		}
		ns[i], ns[m] = n, ns[i]
	}

	return ns
}
