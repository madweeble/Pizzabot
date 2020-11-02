package utils

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

func GetCoords(inputString []string) [][]int {
	exampleCoords := make([][]int, 2)
	exampleCoords[0] = []int{1,3}
	exampleCoords[1] = []int{4,4}
	return exampleCoords
}

func getCord(str string) []int {
	if !regexp.MustCompile("\\(\\s*[0-9]+,\\s*[0-9]+\\)").MatchString(str) {
		log.Println("error occurred while parsing co-ordinate")
		//log.Fatal(errors.New("error occurred while parsing co-ordinate"))
		return nil
	}

	xyCoord := regexp.MustCompile("[0-9]+")
	if !xyCoord.MatchString(str) {
		log.Println("error occurred while parsing co-ordinate")
		//log.Fatal(errors.New("error occurred while parsing co-ordinate"))
		return nil
	}
	result := xyCoord.FindAllString(str,  2)
	x, _ := strconv.Atoi(result[0])
	y, _ := strconv.Atoi(result[1])
	return []int{x,y}
}

func validateArgs(coordArgs []string) bool {
	validRegex := "([0-9]x[0-9])\\s*(\\(\\s*[0-9]+,\\s*[0-9]+\\)\\s*)+"
	validCoordinates := regexp.MustCompile(validRegex)
	argsString := strings.Join(coordArgs, "")

	if validCoordinates.MatchString(argsString) {
		return true
	}

	log.Println("error occurred while parsing arguments")
	//log.Fatal(errors.New("error occurred while parsing co-ordinate"))
	return false
}