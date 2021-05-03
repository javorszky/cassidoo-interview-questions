package may032021

import "testing"

func Test_round(t *testing.T) {
	type args struct {
		n float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "rounds 1.7",
			args: args{n: 1.7},
			want: 1,
		},
		{
			name: "rounds -2.1",
			args: args{n: -2.1},
			want: -2,
		},
		{
			name: "rounds 500.4",
			args: args{n: 500.4},
			want: 500,
		},
		{
			name: "rounds -369.5",
			args: args{n: -369.5},
			want: -369,
		},
		{
			name: "rounds 150",
			args: args{n: 150},
			want: 150,
		},
		{
			name: "rounds -350",
			args: args{n: -350},
			want: -350,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := round(tt.args.n); got != tt.want {
				t.Errorf("round() = %v, want %v", got, tt.want)
			}
		})
	}
}
