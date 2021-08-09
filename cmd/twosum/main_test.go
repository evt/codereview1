package main

import (
	"fmt"
	"reflect"
	"testing"
)

var cases = []struct {
	data     []int
	target   int
	expected []int
}{
	{
		[]int{1, 2, 3},
		3,
		[]int{0, 1},
	},
	{
		[]int{2, 3, 4, 5},
		8,
		[]int{1, 3},
	},
	{
		[]int{3, 3},
		6,
		[]int{0, 1},
	},
	{
		[]int{3, 3},
		8,
		nil,
	},
}

func Test_doubleSumCycle(t *testing.T) {
	for _, cs := range cases {
		t.Run(fmt.Sprintf("Case %v, target %v", cs.data, cs.target), func(t *testing.T) {
			result := doubleSumCycle(cs.data, cs.target)
			if !reflect.DeepEqual(result, cs.expected) {
				t.Error("unexpected result")
			}
		})
	}
}

func Test_doubleSumMap(t *testing.T) {
	for _, cs := range cases {
		t.Run(fmt.Sprintf("Case %v, target %v", cs.data, cs.target), func(t *testing.T) {
			result := doubleSumMap(cs.data, cs.target)
			if !reflect.DeepEqual(result, cs.expected) {
				t.Error("unexpected result")
			}
		})
	}
}

func Benchmark_doubleSumCycle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doubleSumCycle([]int{2, 4, 5, 7, 11, 3, 6, 13, 10, 32, 23, 22, 44, 33, 1, 14, 42, 17, 26, 80}, 102)
	}
}

func Benchmark_doubleSumMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doubleSumMap([]int{2, 4, 5, 7, 11, 3, 6, 13, 10, 32, 23, 22, 44, 33, 1, 14, 42, 17, 26, 80}, 102)
	}
}
