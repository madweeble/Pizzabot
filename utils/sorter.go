package utils

import "sort"

// Sort will take a 2-dimensional array of ints and sort based firstly
// on the first (x) value, then on the second (y) value for each x value
func Sort(coords [][2]int) [][2]int {
	coords = sortByX(coords)
	coords = sortByY(coords)
	return coords
}

func sortByX(coords [][2]int) [][2]int {
	sort.SliceStable(coords, func(i, j int) bool {
		return coords[i][0] < coords[j][0]
	})
	return coords
}

func sortByY(coords [][2]int) [][2]int {
	// TODO: add code here to sort each x-level co-ordinate by the y value
	return coords
}
