package main

import (
	. "Pizzabot/utils"
	"fmt"
	"os"
)

func main() {
	// get an array of co-ordinates from the input args
	gridSize, coordArray := GetCoords(os.Args[:])
	coordArray = Sort(coordArray)

	// calculate and print the route
	route := traverseRoute(coordArray[0:])
	fmt.Printf("Grid Size: %v, Route: %v\n", gridSize, route)
}

func traverseRoute(coords [][2]int) string {
	route := ""
	curr := [2]int{0,0}

	for i := 0; i < len(coords); i++ {
		next := coords[i]
		x := next[0] - curr[0]
		y := next[1] - curr[1]
		route += getNextRoute(x,y) + "D"
		curr = next
	}

	return route
}

func getNextRoute(x int, y int) string {
	route := ""

	if x > 0 {
		route += getDirection("E", x)
	} else if x < 0 {
		route += getDirection("W", -x)
	}

	if y > 0 {
		route += getDirection("N", y)
	} else if y < 0 {
		route += getDirection("S", -y)
	}

	return route
}

func getDirection(d string, count int) string {
	direction := ""

	for i := 0; i < count; i++ {
		direction += d
	}

	return direction
}