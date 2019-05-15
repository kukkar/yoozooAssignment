package main

import (
	"fmt"
	"yoozooAssignment/minimizeFair/travelCost"
)

var people = []string{"A", "B", "C"}
var destinations = []string{"D1", "D2", "D3"}

func main() {
	soln := travelCost.TravelCost(people, destinations)
	fmt.Println(soln)
}
