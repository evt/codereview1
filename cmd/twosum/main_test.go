package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"sync"
	"testing"
	"time"
)

type TestData struct {
	data   []int
	target int
}

var (
	data TestData

	cases = []struct {
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
)

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

func getTestData() *TestData {
	once := &sync.Once{}
	once.Do(func() {
		sums := make(map[int]bool)
		var target int
		r := rand.New(rand.NewSource(time.Now().Unix()))
		nums := r.Perm(2000)
		for i := 0; i < len(nums)-1; i++ {
			for j := i + 1; j < len(nums); j++ {
				if _, exist := sums[nums[i]+nums[j]]; exist {
					sums[nums[i]+nums[j]] = false
				} else {
					sums[nums[i]+nums[j]] = true
				}
			}
		}
		for key, val := range sums {
			if val {
				target = key
				break
			}
		}

		data.data = nums
		data.target = target
	})
	return &data
}

func Benchmark_doubleSumCycle(b *testing.B) {
	data := getTestData()
	for i := 0; i < b.N; i++ {
		doubleSumCycle(data.data, data.target)
	}
}

func Benchmark_doubleSumMap(b *testing.B) {
	data := getTestData()
	for i := 0; i < b.N; i++ {
		doubleSumMap(data.data, data.target)
	}
}
