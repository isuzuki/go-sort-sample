package main

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	ns := []int{5, 2, 1, 1, 7}
	r := sort(ns)

	if !reflect.DeepEqual(r, []int{1, 1, 2, 5, 7}) {
		t.Fatalf("failed test %v", r)
	}
}
