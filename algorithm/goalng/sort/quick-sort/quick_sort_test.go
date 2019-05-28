package qsort

import (
	"math/rand"
	"testing"
)

func generateRandomArr(length int) []int {
	arr := make([]int, length, length)
	for i := 0; i < length; i++ {
		arr[i] = rand.Intn(100)
	}
	return arr
}

func testQuickSort1(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test quick sort 1",
			args: args{
				arr: []int{48, 83, 58, 43, 45, 88}, // generateRandomArr(10),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QuickSort1(tt.args.arr)
			t.Logf("result: %v\n", tt.args.arr)
		})
	}
}

func testQuickSort2(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test quick sort 1",
			args: args{
				arr: []int{48, 83, 58, 43, 45, 88}, // generateRandomArr(10),
			},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QuickSort2(tt.args.arr)
			t.Logf("result: %v\n", tt.args.arr)
		})
	}
}

func BenchmarkQuickSort3(t *testing.B) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test quick sort 1",
			args: args{
				arr: []int{48, 83, 58, 43, 45, 88}, // generateRandomArr(10),
			},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			QuickSort3(tt.args.arr)
			t.Logf("result: %v\n", tt.args.arr)
		})
	}
}
