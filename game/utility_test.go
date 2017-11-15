package game

import (
	"testing"

	"github.com/mitchellh/hashstructure"
)

func TestCoordinateOverlaps(t *testing.T) {
	allCoords := []Coordinate{
		Coordinate{
			X: 1,
			Y: 1,
		},
		Coordinate{
			X: 2,
			Y: 1,
		},
		Coordinate{
			X: 2,
			Y: 1,
		},
	}
	coord := Coordinate{
		X: 2,
		Y: 1,
	}

	if !coordinateOverlaps(coord, allCoords) {
		t.Error("Expected coordinates to overlap")
	}
}

func TestCoordinateWithinBounds(t *testing.T) {
	board := Board{
		RowSize:    4,
		ColumnSize: 6,
	}

	validCoord := Coordinate{
		X: 2,
		Y: 1,
	}

	if !coordinateWithinBounds(validCoord, board) {
		t.Error("Expected coordinate to be within bounds")
	}

	invalidCoord1 := Coordinate{
		X: 0,
		Y: 0,
	}

	if coordinateWithinBounds(invalidCoord1, board) {
		t.Error("Expected coordinate to not be within bounds")
	}

	invalidCoord2 := Coordinate{
		X: 3,
		Y: 0,
	}

	if coordinateWithinBounds(invalidCoord2, board) {
		t.Error("Expected coordinate to not be within bounds")
	}
}

func TestPieceValid(t *testing.T) {
	board := Board{
		RowSize:    4,
		ColumnSize: 6,
		Pieces: []Piece{
			Piece{
				Coordinates: []Coordinate{
					Coordinate{
						X: 2,
						Y: 2,
					},
				},
			},
		},
	}

	hash, _ := hashstructure.Hash(board, nil)
	existingHashes := []uint64{hash}

	board.Pieces[0].Coordinates[0].X = 3

	if !pieceValid(board.Pieces[0], board, existingHashes) {
		t.Error("Expected piece to be valid")
	}

	board.Pieces[0].Coordinates[0].X = 2

	if pieceValid(board.Pieces[0], board, existingHashes) {
		t.Error("Expected piece to not be valid")
	}
}
