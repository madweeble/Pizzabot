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
		{"One north", args{"N",1}, "N"},
		{"One south", args{"S",1}, "S"},
		{"One east", args{"E",1}, "E"},
		{"One west", args{"W",1}, "W"},
		{"Three north", args{"N",3}, "NNN"},
		{"Five east", args{"E",5}, "EEEEE"},
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
		{"1 x 1", args{1,1}, "NE"},
		{"-1 x -1", args{-1,-1}, "SW"},
		{"-3 x 5", args{-3,5}, "SSSEEEEE"},
		{"3 x -5", args{3,-5}, "NNNWWWWW"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNextRoute(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("getRoute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_traverseRoute(t *testing.T) {
	type args struct {
		coords [][2]int
	}
	arr1 := [][2]int{{1,3},{5,5}}
	arr2 := [][2]int{{1,2},{2,2},{4,3},{5,5}}
	arr3 := [][2]int{{1}}
	arr4 := [][2]int{{}}
	tests := []struct {
		name    string
		args    args
		want    string
	}{
		{"2 locations", args{arr1}, "NEEEDNNNNEED"},
		{"4 locations",args{arr2}, "NEEDNDNNEDNEED"},
		{"1 location",args{arr3}, "ND"},
		{"0 locations",args{arr4}, "D"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := traverseRoute(tt.args.coords); got != tt.want {
				t.Errorf("traverseRoute() got = %v, want %v", got, tt.want)
			}
		})
	}
}
