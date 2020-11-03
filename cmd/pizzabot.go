package main

import (
	. "Pizzabot/utils"
	"errors"
	"fmt"
	"os"
)

func main() {
	gridSize := os.Args[1]
	coords := os.Args[2:]
	fmt.Println("1. Grid size:", gridSize)
	fmt.Println("2. Co-ordinates:", coords)
	coordArray := GetCoords(coords)
	fmt.Println("3. getCords:", coordArray)

	route, _ := traverseRoute(coordArray)
	fmt.Println("4: Route:", route)
}

func traverseRoute(coords [][]int) (string, error) {
	route := ""
	curr := []int{0,0}

	for i := 0; i < len(coords); i++ {
		next := coords[i]
		if len(next) != 2 {
			return "", errors.New("unexpected error while parsing co-ordinate array")
		}
		x := next[0] - curr[0]
		y := next[1] - curr[1]
		route += getRoute(x,y) + "D"
		curr = next
	}

	return route, nil
}

func getRoute(x int, y int) string {
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