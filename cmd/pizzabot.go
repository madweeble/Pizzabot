package main

import (
	. "Pizzabot/utils"
	"fmt"
	"os"
)

func main() {
	// get an array of co-ordinates from the input args
	gridSize, coordArray := GetCoords(os.Args[:])
	fmt.Println("Grid Size:", gridSize)
	fmt.Println("Co-ordinates:", coordArray)
	fmt.Println()

	coordArray = Sort(coordArray)
	fmt.Println("Sorted:", coordArray)

	// calculate and print the route
	route := traverseRoute(coordArray[0:])
	fmt.Println("Route:", route)
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
		route += getDirection("N", x)
	} else if x < 0 {
		route += getDirection("S", -x)
	}

	if y > 0 {
		route += getDirection("E", y)
	} else if y < 0 {
		route += getDirection("W", -y)
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