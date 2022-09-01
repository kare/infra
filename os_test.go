package infra_test

import (
	"os"
	"strings"
	"testing"

	"kkn.fi/infra"
)

type env []string

func setenv(t *testing.T, env env) {
	for _, e := range env {
		vars := strings.Split(e, "=")
		if len(vars) != 2 {
			t.Errorf("expecting a key value pair separated by equal sign, got '%v'", vars)
		}
		key := vars[0]
		value := vars[1]
		os.Setenv(key, value)
	}
}

func TestGetenvDefault(t *testing.T) {
	type args struct {
		key          string
		defaultValue string
	}
	tests := []struct {
		name string
		env  env
		args args
		want string
	}{
		{
			name: "key not found",
			env: env{
				"TACO=0x0000",
			},
			args: args{
				key:          "PIZZA",
				defaultValue: "0xbeef",
			},
			want: "0xbeef",
		},
		{
			name: "key set, but empty",
			env: env{
				"TACO=",
			},
			args: args{
				key:          "TACO",
				defaultValue: "0xbeef",
			},
			want: "0xbeef",
		},
		{
			name: "key found",
			env: env{
				"TACO=0xbeef",
			},
			args: args{
				key:          "TACO",
				defaultValue: "0x0000",
			},
			want: "0xbeef",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			os.Clearenv()
			setenv(t, tt.env)
			got := infra.GetenvDefault(tt.args.key, tt.args.defaultValue)
			if got != tt.want {
				t.Errorf("GetenvDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}
