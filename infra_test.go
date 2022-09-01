package infra_test

import (
	"errors"
	"reflect"
	"testing"

	"kkn.fi/infra"
)

func TestMust(t *testing.T) {
	type args struct {
		value any
		err   error
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			name: "returns value, no panic",
			args: args{
				value: "abc",
				err:   nil,
			},
			want: "abc",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := infra.Must(tt.args.value, tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Must() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMustPanicOnError(t *testing.T) {
	type args struct {
		value any
		err   error
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			name: "panic on error",
			args: args{
				value: "",
				err:   errors.New("error that causes panic"),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		tc := tt
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			defer func() {
				if err := recover(); err == nil {
					t.Error("expecting infra.Must() to panic on error")
				}
			}()
			_ = infra.Must(tc.args.value, tc.args.err)
		})
	}
}
