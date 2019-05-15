package travelCost

import "fmt"

//
//find the travel cost in recursive manner
//
func TravelCost(people, destinations []string) Solution {

	if len(people) == 1 && len(destinations) == 1 {
		soln := Solution{
			Cost: CostMatrix[fmt.Sprintf("%s -> %s", people[0], destinations[0])],
			Path: fmt.Sprintf("%s -> %s", people[0], destinations[0]),
		}
		return soln
	}

	//make sure we are not modifying the received array.
	tmpPeople := make([]string, len(people))
	copy(tmpPeople, people)
	//pick one person
	person := tmpPeople[0]
	tmpPeople = tmpPeople[1:]

	// solutions := make([]Solution, 0, len(destinations))
	var minSoln Solution

	//check for all destinations
	for k, destination := range destinations {

		//cost of travelling chosen person to chosen destination
		c := CostMatrix[fmt.Sprintf("%s -> %s", person, destination)]
		path := fmt.Sprintf("%s -> %s", person, destination)

		soln := Solution{
			Cost: c, Path: path,
		}
		tmpDestination := make([]string, len(destinations))
		copy(tmpDestination, destinations)
		tmpDestination = append(tmpDestination[:k], tmpDestination[k+1:]...)

		recvdSoln := TravelCost(tmpPeople, tmpDestination)

		soln.Cost = soln.Cost + recvdSoln.Cost
		soln.Path = fmt.Sprintf("%s | %s", soln.Path, recvdSoln.Path)

		if soln.Cost < minSoln.Cost || minSoln.Cost == 0 {
			minSoln = soln
		}

	}

	return minSoln
}
