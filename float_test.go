package main

import "testing"

func TestFloat64Equal(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "equal",
			args: args{1.3, 1.3},
			want: true,
		},
		{
			name: "y too big",
			args: args{1.3, 1.3 + 2 * floatEpsilon},
			want: false,
		},
		{
			name: "y too small",
			args: args{1.3, 1.3 + 2 * floatEpsilon},
			want: false,
		},
		{
			name: "y a little big",
			args: args{1.3, 1.3 + floatEpsilon / 10},
			want: true,
		},
		{
			name: "y a little small",
			args: args{1.3, 1.3 + floatEpsilon / 10},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float64Equal(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Float64Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}
