package tests

import "testing"

func TestCompareUnorderedSlice(t *testing.T) {
	type args struct {
		x []int
		y []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "equal - the same order",
			args: args{
				x: []int{1, 2, 3},
				y: []int{1, 2, 3},
			},
			want: true,
		},
		{
			name: "equal - mixed order",
			args: args{
				x: []int{1, 2, 3},
				y: []int{3, 2, 1},
			},
			want: true,
		},
		{
			name: "not equal - wrong len",
			args: args{
				x: []int{1, 2, 3},
				y: []int{3, 2},
			},
			want: false,
		},
		{
			name: "not equal - different elements",
			args: args{
				x: []int{1, 2, 3},
				y: []int{1, 2, 4},
			},
			want: false,
		},
		{
			name: "not equal - different elements",
			args: args{
				x: []int{20, 20, 20, 20, 40},
				y: []int{20, 20, 20, 20, 20},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareUnorderedSlice(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("CompareUnorderedSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
