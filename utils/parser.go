package utils

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var (
	gridX = 0
	gridY = 0
)

// GetCoords will accept the input arguments and
// return the grid size and an array of x,y co-ordinates
func GetCoords(inputString []string) (string, [][2]int) {
	// convert input args to string and validate
	argsString := strings.Join(inputString[1:], "")
	if !validateArgs(argsString) {
		log.Fatal(errors.New(fmt.Sprintf("ERROR: incorrect arguments, expected: "+
			"'grid_size (co-ordinates_1) ... (co-ordinates_n)', got: %v", argsString)))
	}

	args := strings.SplitN(argsString, " ", 2)
	gridX, gridY = setGrid(args[0])
	coordStr := strings.Join(args[1:], "")
	coordArray := getCoords(coordStr)

	return args[0], coordArray
}

func setGrid(grid string) (int, int) {
	sizes := strings.SplitN(grid, "x", 2)
	x, err := strconv.Atoi(sizes[0])
	if err != nil {
		log.Fatalf("ERROR: failed to parse '%v' into grid size for 'x'", sizes[0])
	}
	y, err := strconv.Atoi(sizes[1])
	if err != nil {
		log.Fatalf("ERROR: failed to parse '%v' into grid size for 'y'", sizes[1])
	}
	return x, y
}

func getCoords(argsString string) [][2]int {
	// get a collection of numbers and start putting them into arrays
	newCoords := strings.FieldsFunc(argsString, func(r rune) bool {
		return r == '(' || r == ',' || r == ' ' || r == ')'
	})
	numOfCoords := len(newCoords)

	if numOfCoords%2 != 0 {
		log.Fatalf("we require a matching number of x and y co-ordinates,"+
			" but there were %v total co-ordinates: %v", numOfCoords, newCoords)
	}

	locations := make([][2]int, numOfCoords/2)
	arrayCounter := 0
	for nextLocation := 0; nextLocation < numOfCoords/2; nextLocation++ {
		x, _ := strconv.Atoi(newCoords[arrayCounter])
		if x > gridX {
			log.Fatalf("ERROR: 'x'-value [%v] is outside the expected range [%v]", x, gridX)
		}
		locations[nextLocation][0] = x
		arrayCounter++
		y, _ := strconv.Atoi(newCoords[arrayCounter])
		if y > gridY {
			log.Fatalf("ERROR: 'y'-value [%v] is outside the expected range [%v]", y, gridY)
		}
		locations[nextLocation][1] = y
		arrayCounter++
	}

	return locations
}

func validateArgs(argsString string) bool {
	validRegex := "(^[0-9]+x[0-9]+)\\s*(\\(\\s*[0-9]+,\\s*[0-9]+\\)\\s*)+"
	validCoordinates := regexp.MustCompile(validRegex)
	return validCoordinates.MatchString(argsString)
}
