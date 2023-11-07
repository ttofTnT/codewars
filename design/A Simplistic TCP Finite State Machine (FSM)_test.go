package design

import "testing"

func TestTraverseTCPStates(t *testing.T) {
	type args struct {
		events []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"text1",
			args{
				[]string{"APP_ACTIVE_OPEN", "RCV_SYN_ACK", "RCV_FIN"},
			},
			"CLOSE_WAIT",
		},
		{
			"text2",
			args{
				[]string{"APP_PASSIVE_OPEN", "RCV_SYN", "RCV_ACK"},
			},
			"ESTABLISHED",
		},
		{
			"text3",
			args{
				[]string{"APP_ACTIVE_OPEN", "RCV_SYN_ACK", "RCV_FIN", "APP_CLOSE"},
			},
			"LAST_ACK",
		},
		{
			"text4",
			args{
				[]string{"APP_ACTIVE_OPEN"},
			},
			"SYN_SENT",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TraverseTCPStates(tt.args.events); got != tt.want {
				t.Errorf("TraverseTCPStates() = %v, want %v", got, tt.want)
			}
		})
	}
}
