package game

func filterPieces(pieces []Piece, f func(Piece) bool) []Piece {
	piecesf := make([]Piece, 0)
	for _, piece := range pieces {
		if f(piece) {
			piecesf = append(piecesf, piece)
		}
	}
	return piecesf
}

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

func mapCoordinates(coordinates []Coordinate, f func(Coordinate) Coordinate) []Coordinate {
	coordinatesm := make([]Coordinate, len(coordinates))
	for i, v := range coordinates {
		coordinatesm[i] = f(v)
	}
	return coordinatesm
}

func hashIndex(hashes []uint64, hash uint64) int {
	for i, v := range hashes {
		if v == hash {
			return i
		}
	}
	return -1
}
func includesHash(hashes []uint64, hash uint64) bool {
	return hashIndex(hashes, hash) >= 0
}
