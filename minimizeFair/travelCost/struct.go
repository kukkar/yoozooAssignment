package travelCost

import "fmt"

type Solution struct {
	Cost int
	Path string
}

var CostMatrix map[string]int = map[string]int{
	"A -> D1": 2500,
	"A -> D2": 4000,
	"A -> D3": 3500,

	"B -> D1": 4000,
	"B -> D2": 6000,
	"B -> D3": 3500,

	"C -> D1": 2000,
	"C -> D2": 4000,
	"C -> D3": 2500,
}

func (this Solution) String() string {
	return fmt.Sprintf("Cost: %d, Path: %s", this.Cost, this.Path)
}
