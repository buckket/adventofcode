package helper

import "sort"

func AbsInt(x int) int {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}

func Median(states []int) int {
	sort.Ints(states)
	if len(states)%2 != 0 {
		return states[len(states)/2]
	} else {
		return (states[len(states)/2-1] + states[len(states)/2+1]) / 2
	}
}

func Mean(states []int) int {
	var sum int
	for _, v := range states {
		sum += v
	}
	return sum / len(states)
}
