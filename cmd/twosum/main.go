package main

import (
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	return nil
}

func doubleSumCycle(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func doubleSumMap(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for idx, x := range nums {
		if index, founded := m[target-x]; founded {
			return []int{index, idx}
		}
		m[x] = idx
	}
	return nil
}
