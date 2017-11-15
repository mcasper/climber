package game

import (
	// "fmt"

	"github.com/mitchellh/hashstructure"
)

func atFinishLine(piece Piece) bool {
	return anyCoordinate(piece.Coordinates, func(c Coordinate) bool { return c.X == 1 && c.Y == 0 }) &&
		anyCoordinate(piece.Coordinates, func(c Coordinate) bool { return c.X == 2 && c.Y == 0 })
}

func isSolution(board Board) bool {
	for _, piece := range board.Pieces {
		if atFinishLine(piece) {
			return true
		}
	}

	return false
}

func makeMoves(board Board, allBoardHashes []uint64) []Board {
	boards := make([]Board, 0)

	for i := 0; i < len(board.Pieces); i++ {
		newBoardXLeft := cloneBoard(board)
		newBoardXLeft.Layer = board.Layer + 1
		pieceXLeft := newBoardXLeft.Pieces[i]
		for i2 := 0; i2 < len(pieceXLeft.Coordinates); i2++ {
			newBoardXLeft.Pieces[i].Coordinates[i2].X = newBoardXLeft.Pieces[i].Coordinates[i2].X - 1
		}

		if pieceValid(pieceXLeft, newBoardXLeft, allBoardHashes) {
			boards = append(boards, newBoardXLeft)
		}

		newBoardXRight := cloneBoard(board)
		newBoardXRight.Layer = board.Layer + 1
		pieceXRight := newBoardXRight.Pieces[i]
		for i2 := 0; i2 < len(pieceXRight.Coordinates); i2++ {
			newBoardXRight.Pieces[i].Coordinates[i2].X = newBoardXRight.Pieces[i].Coordinates[i2].X + 1
		}

		if pieceValid(pieceXRight, newBoardXRight, allBoardHashes) {
			boards = append(boards, newBoardXRight)
		}

		newBoardYUp := cloneBoard(board)
		newBoardYUp.Layer = board.Layer + 1
		pieceYUp := newBoardYUp.Pieces[i]
		for i2 := 0; i2 < len(pieceYUp.Coordinates); i2++ {
			newBoardYUp.Pieces[i].Coordinates[i2].Y = newBoardYUp.Pieces[i].Coordinates[i2].Y - 1
		}

		if pieceValid(pieceYUp, newBoardYUp, allBoardHashes) {
			boards = append(boards, newBoardYUp)
		}

		newBoardYDown := cloneBoard(board)
		newBoardYDown.Layer = board.Layer + 1
		pieceYDown := newBoardYDown.Pieces[i]
		for i2 := 0; i2 < len(pieceYDown.Coordinates); i2++ {
			newBoardYDown.Pieces[i].Coordinates[i2].Y = newBoardYDown.Pieces[i].Coordinates[i2].Y + 1
		}

		if pieceValid(pieceYDown, newBoardYDown, allBoardHashes) {
			boards = append(boards, newBoardYDown)
		}
	}

	return boards
}

func pieceValid(piece Piece, board Board, allBoardHashes []uint64) bool {
	hash, _ := hashstructure.Hash(board, nil)
	if includesHash(allBoardHashes, hash) {
		return false
	}

	totalCoordinates := make([]Coordinate, 0)
	for _, piece := range board.Pieces {
		for _, coordinate := range piece.Coordinates {
			totalCoordinates = append(totalCoordinates, coordinate)
		}
	}

	return allCoordinates(piece.Coordinates, func(c Coordinate) bool {
		return coordinateValid(c, totalCoordinates, board)
	})
}

func coordinateValid(c Coordinate, totalCoordinates []Coordinate, board Board) bool {
	return coordinateWithinBounds(c, board) && !coordinateOverlaps(c, totalCoordinates)
}

func coordinateWithinBounds(c Coordinate, board Board) bool {
	return c.X >= 0 &&
		c.X <= (board.RowSize-1) &&
		c.Y >= 0 &&
		c.Y <= (board.ColumnSize-1) &&
		!(c.X == 0 && c.Y == 0) &&
		!(c.X == 3 && c.Y == 0)
}

func coordinateOverlaps(c Coordinate, totalCoordinates []Coordinate) bool {
	return len(filterCoordinates(totalCoordinates, func(c2 Coordinate) bool {
		return c.X == c2.X && c.Y == c2.Y
	})) > 1
}
