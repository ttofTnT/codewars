package backtracking

import (
	"reflect"
	"testing"
)

func TestPermutations(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"test1", args{"abc"}, []string{"abc", "acb", "bac", "bca", "cab", "cba"}},
		{"test2", args{"ab"}, []string{"ab", "ba"}},
		{"test3", args{"a"}, []string{"a"}},
		{"test4", args{"aabb"}, []string{"aabb", "abab", "abba", "baab", "baba", "bbaa"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Permutations(tt.args.s); !reflect.DeepEqual(got, tt.want) {

				t.Errorf("Permutations() = %v, want %v", got, tt.want)
			}
		})
	}
}
