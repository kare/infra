package env_test

import (
	"testing"

	"kkn.fi/infra/env"
)

func TestParseUnknownEnv(t *testing.T) {
	env, err := env.ParseEnv("x")
	if err == nil {
		t.Errorf("expecting error, got nil")
	}
	if env.String() != "" {
		t.Errorf("expecting empty Env")
	}
}

func TestParseDevelopmentEnv(t *testing.T) {
	e, err := env.ParseEnv("development")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if e != env.Development {
		t.Errorf("expecting env '%v', got '%v'", "development", e)
	}
}

func TestParseEnv(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    env.Env
		wantErr bool
	}{
		{
			name: "development",
			args: args{
				input: "development",
			},
			want:    env.Development,
			wantErr: false,
		},
		{
			name: "uppercase development",
			args: args{
				input: "DEVELOPMENT",
			},
			want:    env.Development,
			wantErr: false,
		},
		{
			name: "staging",
			args: args{
				input: "staging",
			},
			want:    env.Staging,
			wantErr: false,
		},
		{
			name: "mixed case staging",
			args: args{
				input: "sTaGinG",
			},
			want:    env.Staging,
			wantErr: false,
		},
		{
			name: "production",
			args: args{
				input: "production",
			},
			want:    env.Production,
			wantErr: false,
		},
		{
			name: "upper first production",
			args: args{
				input: "Production",
			},
			want:    env.Production,
			wantErr: false,
		},
		{
			name: "unknown",
			args: args{
				input: "unknown",
			},
			want:    env.Env(""),
			wantErr: true,
		},
		{
			name: "empty",
			args: args{
				input: "",
			},
			want:    env.Env(""),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			env, err := env.ParseEnv(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("env.ParseEnv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if env != tt.want {
				t.Errorf("expecting env '%v', got '%v'", tt.want.String(), env)
			}
		})
	}
}
