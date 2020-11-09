package main

import "testing"

func Test_traverseRoute(t *testing.T) {
	type args struct {
		coords [][2]int
	}
	arr1 := [][2]int{{1,3},{4,4}}
	arr2 := [][2]int{{1,2},{2,2},{4,3},{5,5}}
	arr3 := [][2]int{{1}}
	arr4 := [][2]int{{}}
	tests := []struct {
		name    string
		args    args
		want    string
	}{
		{"2Locations", args{arr1}, "ENNNDEEEND"},
		{"4Locations",args{arr2}, "ENNDEDEENDENND"},
		{"1Location",args{arr3}, "ED"},
		{"0Locations",args{arr4}, "D"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := traverseRoute(tt.args.coords); got != tt.want {
				t.Errorf("traverseRoute() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getNextRoute(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"0x0", args{0,0}, ""},
		{"1x1", args{1,1}, "EN"},
		{"-1x-1", args{-1,-1}, "WS"},
		{"-3x5", args{-3,5}, "WWWNNNNN"},
		{"3x-5", args{3,-5}, "EEESSSSS"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNextRoute(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("getRoute() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
		{"OneNorth", args{"N",1}, "N"},
		{"OneSouth", args{"S",1}, "S"},
		{"OneEast", args{"E",1}, "E"},
		{"OneWest", args{"W",1}, "W"},
		{"ThreeNorth", args{"N",3}, "NNN"},
		{"FiveEast", args{"E",5}, "EEEEE"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDirection(tt.args.d, tt.args.count); got != tt.want {
				t.Errorf("getDirection() = %v, want %v", got, tt.want)
			}
		})
	}
}
