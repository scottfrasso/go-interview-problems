package main

//"Your car starts at position 0 and speed +1 on an infinite number line.
// (Your car can go into negative positions.) Your car drives automatically
// according to a sequence of instructions A (accelerate) and R (reverse)...
//Now for some target position, say the length of the shortest sequence of instructions to get there."

func racecar(target int) int {
	path := make([]int,0)
	fastestPath := followPath(target, path)

	return len(fastestPath)
}

func followPath(target int, path []int) []int {
	currentNode := 0
	if len(path) != 0 {
		currentNode = path[len(path)-1]
	}
	currentSum := sum(path)

	if currentSum == target {
		return path
	}

	// check if +1 does not go over
	if currentSum + currentNode + 1 <= target {
		path = append(path, currentNode + 1)

		return followPath(target, path)
	}

	// check if staying does not goes over
	if currentSum + currentNode <= target {
		path = append(path, currentNode)

		return followPath(target, path)
	}

	path = append(path, currentNode -1)

	return followPath(target, path)
}

func sum(path []int) int{
	result := 0
	for _, v := range path {
		result += v
	}

	return result
}
