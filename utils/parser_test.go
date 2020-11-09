package utils

import (
	"reflect"
	"testing"
)

func TestGetCoords(t *testing.T) {
	type args struct {
		inputString []string
	}
	args1 := []string{"./pizzabot", "5x5 ", "(1, 3)"}
	args2 := []string{"./pizzabot", "5x5 ", "(1, 3)", "(4, 4)"}
	args3 := []string{"./pizzabot", "15x10 ", "(1, 3)", "(4, 4)", "(12, 1)"}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 [][2]int
	}{
		{"Test1", args{args1}, "5x5", [][2]int{{1,3}}},
		{"Test2", args{args2}, "5x5", [][2]int{{1,3}, {4,4}}},
		{"Test3", args{args3}, "15x10", [][2]int{{1,3}, {4,4}, {12, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetCoords(tt.args.inputString)
			if got != tt.want {
				t.Errorf("GetCoords() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetCoords() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_setGrid(t *testing.T) {
	type args struct {
		grid string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{"0x0", args{"0x0"}, 0, 0},
		{"5x1", args{"5x1"}, 5, 1},
		{"1x5", args{"1x5"}, 1, 5},
		{"5x5", args{"5x5"}, 5, 5},
		{"10x15", args{"10x15"}, 10, 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := setGrid(tt.args.grid)
			if got != tt.want {
				t.Errorf("setGrid() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("setGrid() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_getCoords(t *testing.T) {
	type args struct {
		argsString string
	}
	expect1 := [][2]int{{0, 0}}
	expect2 := [][2]int{{1, 3}, {4, 4}}
	expect3 := [][2]int{{3, 1}, {4, 4}, {12, 7}}
	tests := []struct {
		name    string
		args    args
		want    [][2]int
	}{
		{"Valid x 1", args{"(0, 0)"}, expect1},
		{"Valid x 2", args{"(1, 3), (4, 4)"}, expect2},
		{"Valid x 3", args{"(3, 1), (4, 4), (12, 7)"}, expect3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gridX, gridY = 12, 12
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
		{"Valid 1 location", args{"3x3 (3,1)"}, true},
		{"Invalid location", args{"1x1(3)"}, false},
		{"Invalid grid size", args{"5(1, 3)"}, false},
		{"Negative grid", args{"5x-5 (1, 3)"}, false},
		{"Negative x", args{"5x5 (-1, 3)"}, false},
		{"Negative y", args{"5x5 (1, -3)"}, false},
		{"Negative xy", args{"5x5 (-1, -3)"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateArgs(tt.args.argsString); got != tt.want {
				t.Errorf("validateArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
