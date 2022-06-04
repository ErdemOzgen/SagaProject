package utils

import (
	"testing"
)

func TestSetLogLevel(t *testing.T) {
	type args struct {
		env string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"production test", args{env: "production"}},
		{"development test", args{env: "development"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetLogLevel(tt.args.env)
			t.Log(tt.args.env)
		})
	}
}
