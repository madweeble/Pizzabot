package main

import "testing"

func Test_getDirection(t *testing.T) {
	type args struct {
		d     string
		count int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDirection(tt.args.d, tt.args.count); got != tt.want {
				t.Errorf("getDirection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getRoute(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRoute(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("getRoute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_traverseRoute(t *testing.T) {
	type args struct {
		coords [][]int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := traverseRoute(tt.args.coords)
			if (err != nil) != tt.wantErr {
				t.Errorf("traverseRoute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("traverseRoute() got = %v, want %v", got, tt.want)
			}
		})
	}
}
