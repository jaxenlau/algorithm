package rotateimage

import (
	"fmt"
	"testing"
)

func Test_rotate_simple(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name   string
		args   args
		except args
	}{
		{
			name: "Test_rotate_simple",
			args: args{
				matrix: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			except: args{
				matrix: [][]int{
					{7, 4, 1},
					{8, 5, 2},
					{9, 6, 3},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate_simple(tt.args.matrix)
			n := len(tt.args.matrix)
			for i := 0; i < n; i++ {
				for j := 0; j < n; j++ {
					ret := tt.args.matrix[i][j]
					except := tt.except.matrix[i][j]
					fmt.Printf("%d ", except)
					if ret != except {
						t.Errorf("[%d][%d] -> ret(%d) != except(%d)\n", i, j, ret, except)
					}
				}
				fmt.Printf("\n")
			}
		})
	}
}
