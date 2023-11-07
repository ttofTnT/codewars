package string

import "testing"

func TestGetMinBase(t *testing.T) {
	type args struct {
		n uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			"test1",
			args{
				3,
			},
			2,
		},
		{
			"test2",
			args{
				7,
			},
			2,
		},
		{
			"test3",
			args{
				1111,
			},
			10,
		},
		{
			"test4",
			args{
				1000000000000,
			},
			999999999999,
		},
		{
			"test5",
			args{
				1001001,
			},
			1000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMinBase(tt.args.n); got != tt.want {
				t.Errorf("GetMinBase() = %v, want %v", got, tt.want)
			}
		})
	}
}
