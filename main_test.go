package main

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	cases := []struct {
		a, b, want []int
	}{
		{[]int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
		{[]int{}, []int{}, []int{}},
		{[]int{}, []int{0}, []int{0}},
		{[]int{5}, []int{}, []int{5}},
		{[]int{1, 1}, []int{1}, []int{1, 1, 1}},
		{[]int{-3, 7}, []int{-5, 10}, []int{-5, -3, 7, 10}},
	}
	for _, c := range cases {
		got := toSlice(Merge(fromSlice(c.a), fromSlice(c.b)))
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Merge(%v,%v) = %v, want %v", c.a, c.b, got, c.want)
		}
	}
}
