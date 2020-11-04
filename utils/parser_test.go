package utils

import (
	"reflect"
	"testing"
)

func Test_getCoords(t *testing.T) {
	type args struct {
		argsString string
	}
	tests := []struct {
		name    string
		args    args
		want    [][2]int
		wantErr bool
	}{
		{"Valid x 1", args{"(0, 0)"}, [][2]int{{0, 0}}, false},
		{"Valid x 2", args{"(1, 3), (4, 4)"}, [][2]int{{1, 3}, {4, 4}}, false},
		{"Valid x 3", args{"(3, 1), (4, 4), (12, 7)"}, [][2]int{{3, 1}, {4, 4}, {12, 7}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCoords(tt.args.argsString); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCoords() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateArgs(t *testing.T) {
	type args struct {
		argsString string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Valid 2 locations", args{"5x5 (1, 3), (4, 4)"}, true},
		{"Valid 1 location", args{"3x3(3,1)"}, true},
		{"Invalid location", args{"1x1(3)"}, false},
		{"Invalid grid size", args{"5(1, 3)"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateArgs(tt.args.argsString); got != tt.want {
				t.Errorf("validateArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
