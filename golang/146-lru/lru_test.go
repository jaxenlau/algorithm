package _46_lru

import (
	"testing"
)

// ["LRUCache","put","put","get","put","put","get"]
// [[2],[2,1],[2,2],[2],[1,1],[4,1],[2]]

func TestLRUCache_Put(t *testing.T) {

	lru := Constructor(2)

	type args struct {
		key   int
		value int
	}
	tests := []struct {
		method string
		args   args
		except int
	}{
		{method: "put", args: args{2, 1}},
		{method: "put", args: args{2, 2}},
		{method: "get", args: args{2, 0}, except: 2},
		{method: "put", args: args{1, 1}},
		{method: "put", args: args{4, 1}},
		{method: "get", args: args{2, 0}, except: -1},
	}
	for _, tt := range tests {
		if tt.method == "put" {
			lru.Put(tt.args.key, tt.args.value)
			t.Logf("put: (%d, %d)\n", tt.args.key, tt.args.value)
		} else {
			ret := lru.Get(tt.args.key)
			t.Logf("get value: %#v\n", ret)
			if ret != tt.except {
				t.Errorf("ret: %d, except: %d\n", ret, tt.except)
			}
		}
	}
}
