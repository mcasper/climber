package game

import ()

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

func makeMoves(board Board, allBoardHashes map[string]int) ([]Board, map[string]int) {
	boards := make([]Board, 0)

	for i := 0; i < len(board.Pieces); i++ {
		// Left
	LeftOuterLoop:
		for count := 1; count < 4; count++ {
			newBoardXLeft := cloneBoard(board)
			newBoardXLeft.Layer = board.Layer + 1
			newPieceXLeft := newBoardXLeft.Pieces[i]

			for i2 := 0; i2 < len(newPieceXLeft.Coordinates); i2++ {
				newBoardXLeft.Pieces[i].Coordinates[i2].X = newBoardXLeft.Pieces[i].Coordinates[i2].X - count
				if newBoardXLeft.Pieces[i].Coordinates[i2].X > (board.RowSize-1) || newBoardXLeft.Pieces[i].Coordinates[i2].X < 0 {
					break LeftOuterLoop
				}
			}

			if pieceValid(newPieceXLeft, newBoardXLeft, allBoardHashes) {
				boards = append(boards, newBoardXLeft)
				hash := hashBoard(newBoardXLeft)
				allBoardHashes[hash] = 1
			} else {
				break LeftOuterLoop
			}
		}

		// Right
	RightOuterLoop:
		for count := 1; count < 4; count++ {
			newBoardXRight := cloneBoard(board)
			newBoardXRight.Layer = board.Layer + 1
			newPieceXRight := newBoardXRight.Pieces[i]

			for i2 := 0; i2 < len(newPieceXRight.Coordinates); i2++ {
				newBoardXRight.Pieces[i].Coordinates[i2].X = newBoardXRight.Pieces[i].Coordinates[i2].X + count
				if newBoardXRight.Pieces[i].Coordinates[i2].X > (board.RowSize-1) || newBoardXRight.Pieces[i].Coordinates[i2].X < 0 {
					break RightOuterLoop
				}
			}

			if pieceValid(newPieceXRight, newBoardXRight, allBoardHashes) {
				boards = append(boards, newBoardXRight)
				hash := hashBoard(newBoardXRight)
				allBoardHashes[hash] = 1
			} else {
				break RightOuterLoop
			}
		}

		// Up
	UpOuterLoop:
		for count := 1; count < 6; count++ {
			newBoardYUp := cloneBoard(board)
			newBoardYUp.Layer = board.Layer + 1
			newPieceYUp := newBoardYUp.Pieces[i]

			for i2 := 0; i2 < len(newPieceYUp.Coordinates); i2++ {
				newBoardYUp.Pieces[i].Coordinates[i2].Y = newBoardYUp.Pieces[i].Coordinates[i2].Y - count
				if newBoardYUp.Pieces[i].Coordinates[i2].Y > (board.ColumnSize-1) || newBoardYUp.Pieces[i].Coordinates[i2].Y < 0 {
					break UpOuterLoop
				}
			}

			if pieceValid(newPieceYUp, newBoardYUp, allBoardHashes) {
				boards = append(boards, newBoardYUp)
				hash := hashBoard(newBoardYUp)
				allBoardHashes[hash] = 1
			} else {
				break UpOuterLoop
			}
		}

		// Down
	DownOuterLoop:
		for count := 1; count < 6; count++ {
			newBoardYDown := cloneBoard(board)
			newBoardYDown.Layer = board.Layer + 1
			newPieceYDown := newBoardYDown.Pieces[i]

			for i2 := 0; i2 < len(newPieceYDown.Coordinates); i2++ {
				newBoardYDown.Pieces[i].Coordinates[i2].Y = newBoardYDown.Pieces[i].Coordinates[i2].Y + count
				if newBoardYDown.Pieces[i].Coordinates[i2].Y > (board.ColumnSize-1) || newBoardYDown.Pieces[i].Coordinates[i2].Y < 0 {
					break DownOuterLoop
				}
			}

			if pieceValid(newPieceYDown, newBoardYDown, allBoardHashes) {
				boards = append(boards, newBoardYDown)
				hash := hashBoard(newBoardYDown)
				allBoardHashes[hash] = 1
			} else {
				break DownOuterLoop
			}
		}
	}

	return boards, allBoardHashes
}

func pieceValid(piece Piece, board Board, allBoardHashes map[string]int) bool {
	hash := hashBoard(board)
	if allBoardHashes[hash] == 1 {
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
