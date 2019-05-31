package _42_trapping_tain_water

import (
	"testing"
)

func test_trap(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "暴力解法",
			args: args{
				height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trapViolent(tt.args.height); got != tt.want {
				t.Errorf("trapViolent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func test_trapDynamic(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "动态编程解法-1",
			args: args{
				height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			},
			want: 6,
		},
		{
			name: "动态编程解法-2",
			args: args{
				height: []int{2, 0, 2},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trapDynamic(tt.args.height); got != tt.want {
				t.Errorf("trapDynamic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func test_trapStack(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "栈解法-1",
			args: args{
				height: []int{2, 0, 2},
			},
			want: 2,
		},
		{
			name: "栈解法-2",
			args: args{
				height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			},
			want: 6,
		},
		{
			name: "栈解法-3",
			args: args{
				height: []int{0, 1, 4, 2, 5},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trapStack(tt.args.height); got != tt.want {
				t.Errorf("trapStack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_trapTwoPointers(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "双指针法-1",
			args: args{
				height: []int{2, 0, 2},
			},
			want: 2,
		},
		{
			name: "双指针法-2",
			args: args{
				height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			},
			want: 6,
		},
		{
			name: "双指针法-3",
			args: args{
				height: []int{0, 1, 4, 2, 5},
			},
			want: 2,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trapTwoPointers(tt.args.height); got != tt.want {
				t.Errorf("trapTwoPointers() = %v, want %v", got, tt.want)
			}
		})
	}
}
