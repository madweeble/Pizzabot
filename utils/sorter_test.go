package utils

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	type args struct {
		coords [][2]int
	}
	tests := []struct {
		name string
		args args
		want [][2]int
	}{
		{"test 1", args{[][2]int{{3,4},{1,2}}}, [][2]int{{1,2},{3,4}}},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sort(tt.args.coords); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortByX(t *testing.T) {
	type args struct {
		coords [][2]int
	}
	tests := []struct {
		name string
		args args
		want [][2]int
	}{
		{"test 1", args{[][2]int{{3,4},{1,2}}}, [][2]int{{1,2},{3,4}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sort(tt.args.coords); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortByY(t *testing.T) {
	type args struct {
		coords [][2]int
	}
	tests := []struct {
		name string
		args args
		want [][2]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortByY(tt.args.coords); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortByY() = %v, want %v", got, tt.want)
			}
		})
	}
}