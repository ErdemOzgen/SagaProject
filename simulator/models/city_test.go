package models

import (
	"reflect"
	"testing"
)

func TestNewCity(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want City
	}{
		// TODO: Add test cases.
		{"Generate New City", args{name: "city1"}, NewCity("city1")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCity(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCity() = %v, want %v", got, tt.want)
			}

		})
	}
}
