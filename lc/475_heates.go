package main

import (
	"fmt"
	"sort"
)

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)
	standardRadius := 0
	heatersSize := len(heaters)
	heaterPointer := 0
	for _, house := range houses {
		f := 0
		for heaterPointer+1 < heatersSize && diff(heaters[heaterPointer+1], house) < diff(heaters[heaterPointer], house) {
			standardRadius = min(diff(heaters[heaterPointer+1], house), standardRadius)
			heaterPointer += 1
			f = 1
		}
		if f == 0 {
			standardRadius = max(diff(heaters[heaterPointer], house), standardRadius)
		}
		// fmt.Println(house, standardRadius)
	}

	return standardRadius
}

func main() {
	fmt.Println(findRadius([]int{1, 1, 1, 1, 1, 1, 999, 999, 999, 999, 999}, []int{499, 500, 501}))
	fmt.Println(findRadius([]int{1, 1, 1, 1, 1, 1, 999, 999, 999, 999, 999}, []int{499, 500, 501}))
}
