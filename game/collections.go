package game

func anyCoordinate(coordinates []Coordinate, f func(Coordinate) bool) bool {
	for _, coordinate := range coordinates {
		if f(coordinate) {
			return true
		}
	}
	return false
}

func allCoordinates(coordinates []Coordinate, f func(Coordinate) bool) bool {
	for _, coordinate := range coordinates {
		if !f(coordinate) {
			return false
		}
	}
	return true
}

func filterCoordinates(coordinates []Coordinate, f func(Coordinate) bool) []Coordinate {
	coordinatesf := make([]Coordinate, 0)
	for _, v := range coordinates {
		if f(v) {
			coordinatesf = append(coordinatesf, v)
		}
	}
	return coordinatesf
}
