//go:build !integration

package infra_test

import (
	"errors"
	"log"
	"os"
	"reflect"
	"testing"

	"kkn.fi/infra"
)

// Test that standard Logger can be casted to infra.Logger.
var _ infra.Logger = (*log.Logger)(nil)

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

func TestIsDevelopment(t *testing.T) {
	tests := []struct {
		name string
		env  string
		want bool
	}{
		{
			name: "ENV is set to development",
			env:  "development",
			want: true,
		},
		{
			name: "ENV is production",
			env:  "production",
			want: false,
		},
	}
	for _, tt := range tests {
		tc := tt
		t.Run(tc.name, func(t *testing.T) {
			os.Setenv("ENV", tc.env)
			if got := infra.IsDevelopment(); got != tc.want {
				t.Errorf("IsDevelopment() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestIsProduction(t *testing.T) {
	tests := []struct {
		name string
		env  string
		want bool
	}{
		{
			name: "ENV is set to development",
			env:  "development",
			want: false,
		},
		{
			name: "ENV is production",
			env:  "production",
			want: true,
		},
		{
			name: "ENV is empty",
			env:  "",
			want: false,
		},
		{
			name: "ENV is foobar",
			env:  "foobar",
			want: false,
		},
	}
	for _, tt := range tests {
		tc := tt
		t.Run(tc.name, func(t *testing.T) {
			os.Setenv("ENV", tc.env)
			if got := infra.IsProduction(); got != tc.want {
				t.Errorf("IsProduction() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestIsCI(t *testing.T) {
	tests := []struct {
		name string
		env  string
		want bool
	}{
		{
			name: "CI is set to true",
			env:  "true",
			want: true,
		},
		{
			name: "CI is TRUE",
			env:  "TRUE",
			want: false,
		},
		{
			name: "CI is empty",
			env:  "",
			want: false,
		},
		{
			name: "CI is foobar",
			env:  "foobar",
			want: false,
		},
	}
	for _, tt := range tests {
		tc := tt
		t.Run(tc.name, func(t *testing.T) {
			os.Setenv("CI", tc.env)
			if got := infra.IsCI(); got != tc.want {
				t.Errorf("IsCI() = %v, want %v", got, tc.want)
			}
		})
	}
}
