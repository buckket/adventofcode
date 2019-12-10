package main

import (
	"fmt"
)

func main() {
	n := [6]int{2, 6, 5, 2, 7, 5}
	stop := [6]int{7, 8, 1, 5, 8, 4}

	var err error
	for i := 0; i < 10000; i++ {
		n, err = nextNumber(n)
		if err != nil {
			fmt.Printf("nextNumber error: %s", err)
		}
		for j := 0; j < 6; j++ {
			if n[j] > stop[j] {
				fmt.Printf("End reached, %d solutions", i)
				return
			} else if n[j] < stop[j] {
				break
			}
		}
		fmt.Printf("%d: %d\n", i, n)
	}
}

func nextNumber(n [6]int) ([6]int, error) {
	var double int
	for i, _ := range n {
		if i > 0 {
			if n[i] == n[i-1] {
				double++
			}
			if n[i] < n[i-1] {
				for j := i; j < 6; j++ {
					n[j] = n[i-1]
				}
				return n, nil
			}
		}
	}
	for j := 5; j >= 0; j-- {
		if n[j] < 9 {
			n[j]++
			if j == 5 {
				if n[j-1] == (n[j] - 1) {
					double--
				}
				if double > 0 {
					return n, nil
				} else {
					return nextNumber(n)
				}
			}
			return nextNumber(n)
		} else {
			n[j] = 0
		}
	}
	return n, fmt.Errorf("out of max range")
}
