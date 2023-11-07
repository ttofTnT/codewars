package string

import "testing"

func TestFormatDuration(t *testing.T) {
	type args struct {
		seconds int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{1}, "1 second"},
		{"62", args{62}, "1 minute and 2 seconds"},
		{"120", args{120}, "2 minutes"},
		{"3600", args{3600}, "1 hour"},
		{"3662", args{3662}, "1 hour, 1 minute and 2 seconds"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatDuration(tt.args.seconds); got != tt.want {
				t.Errorf("FormatDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}
