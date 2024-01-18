package api

import (
	"reflect"
	"testing"
)

// This test has only happy path as with current struct it's impossible
// to cause failure of this function.
func Test_mustMarshalError(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "happy path",
			args: args{
				message: "response message",
			},
			want: []byte(
				`{"message":"response message"}`,
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mustMarshalError(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mustMarshalError() = %v, want %v", got, tt.want)
			}
		})
	}
}
