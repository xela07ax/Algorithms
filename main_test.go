// go test -bench=.

package main

import (
	"sorting1/sort_numbers"
	"testing"
)

var sortingNums []int
var sortingNumsBig []int

func BenchmarkBubles(b *testing.B) {
	sortingNums = []int{12, 35, 99, 18, 76, 79, 80, 4, 91, 12, 35, 99, 18, 76, 79, 80, 4, 91}
	for i := 0; i < 100000; i++ {
		sort_numbers.SortBuble(sortingNums)
	}
}

func BenchmarkBucket(b *testing.B) {
	sortingNums = []int{12, 35, 99, 18, 76, 79, 80, 4, 91, 12, 35, 99, 18, 76, 79, 80, 4, 91}
	for i := 0; i < 100000; i++ {
		sort_numbers.SortBlock(sortingNums)
	}
}

func BenchmarkMySort(b *testing.B) {
	sortingNums = []int{12, 35, 99, 18, 76, 79, 80, 4, 91, 12, 35, 99, 18, 76, 79, 80, 4, 91}
	for i := 0; i < 100000; i++ {
		sort_numbers.SortMyArtifice(sortingNums)
	}
}

func BenchmarkBublesBig(b *testing.B) {
	sortingNums = []int{12, 35, 99, 18, 76, 79, 80, 4, 91, 12, 35, 99, 18, 76, 79, 80, 4, 91}
	sortingNumsBig = make([]int, 0, len(sortingNums)*9)
	for i := 0; i < 10; i++ {
		sortingNumsBig = append(sortingNumsBig, sortingNums...)
	}
	for i := 0; i < 100000; i++ {
		sort_numbers.SortBuble(sortingNumsBig)
	}
}

func BenchmarkBucketBig(b *testing.B) {
	sortingNums = []int{12, 35, 99, 18, 76, 79, 80, 4, 91, 12, 35, 99, 18, 76, 79, 80, 4, 91}
	sortingNumsBig = make([]int, 0, len(sortingNums)*9)
	for i := 0; i < 10; i++ {
		sortingNumsBig = append(sortingNumsBig, sortingNums...)
	}
	for i := 0; i < 100000; i++ {
		sort_numbers.SortBlock(sortingNumsBig)
	}
}

func BenchmarkMySortBig(b *testing.B) {
	sortingNums = []int{12, 35, 99, 18, 76, 79, 80, 4, 91, 12, 35, 99, 18, 76, 79, 80, 4, 91}
	sortingNumsBig = make([]int, 0, len(sortingNums)*9)
	for i := 0; i < 10; i++ {
		sortingNumsBig = append(sortingNumsBig, sortingNums...)
	}
	for i := 0; i < 100000; i++ {
		sort_numbers.SortMyArtifice(sortingNumsBig)
	}
}
