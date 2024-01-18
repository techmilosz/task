package api

import (
	"testing"

	"repartners/pkg/tests"
)

func Test_transformCalculateResponse(t *testing.T) {
	type args struct {
		vals []int
	}
	tc := []struct {
		name string
		args args
		want []payloadRow
	}{
		{
			name: "happy path - single element",
			args: args{
				vals: []int{10, 10, 10},
			},
			want: []payloadRow{
				{
					Packet: 10,
					Amount: 3,
				},
			},
		},
		{
			name: "happy path - multiple elements",
			args: args{
				vals: []int{10, 20, 10, 30, 10},
			},
			want: []payloadRow{
				{
					Packet: 10,
					Amount: 3,
				},
				{
					Packet: 20,
					Amount: 1,
				},
				{
					Packet: 30,
					Amount: 1,
				},
			},
		},
	}
	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			got := transformCalculateResponse(tt.args.vals)
			if !tests.CompareUnorderedSlice(got, tt.want) {
				t.Errorf("transformCalculateResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
