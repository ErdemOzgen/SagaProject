package utils

import (
	"reflect"
	"testing"
)

func TestGenerateAlienNames(t *testing.T) {
	type args struct {
		count int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{"test 1", args{count: 10}, []string{"Bryant", "Bryant", "Bryant", "Bryant", "Bryant", "Bryant", "Bryant", "Bryant", "Bryant", "Bryant"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateAlienNames(tt.args.count); !reflect.DeepEqual(len(got), len(tt.want)) {
				t.Errorf("GenerateAlienNames() = %v, want %v", got, tt.want)
			}
		})
	}
}
