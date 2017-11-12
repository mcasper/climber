package game

import (
	"fmt"
)

type Coordinate struct {
	X int
	Y int
}

type Piece struct {
	Coordinates []Coordinate
	Shape       string
}

type Board struct {
	Pieces     []Piece
	RowSize    int
	ColumnSize int
}

var TenBoard = Board{
	RowSize:    4,
	ColumnSize: 6,
	Pieces: []Piece{
		Piece{
			Shape: "l",
			Coordinates: []Coordinate{
				Coordinate{
					X: 0,
					Y: 1,
				},
				Coordinate{
					X: 0,
					Y: 2,
				},
				Coordinate{
					X: 1,
					Y: 2,
				},
			},
		},
		Piece{
			Shape: "dot",
			Coordinates: []Coordinate{
				Coordinate{
					X: 2,
					Y: 2,
				},
			},
		},
		Piece{
			Shape: "line",
			Coordinates: []Coordinate{
				Coordinate{
					X: 3,
					Y: 1,
				},
				Coordinate{
					X: 3,
					Y: 2,
				},
			},
		},
	},
}

func PrintBoard(board Board) {
	prettyBoard := make([]string, board.RowSize*board.ColumnSize)

	for i := 0; i < (board.RowSize * board.ColumnSize); i++ {
		prettyBoard[i] = " "
	}

	for i := 0; i < len(board.Pieces); i++ {
		piece := board.Pieces[i]
		for i := 0; i < len(piece.Coordinates); i++ {
			coordinate := piece.Coordinates[i]

			prettyIndex := coordinate.X + (coordinate.Y * 4)
			prettyBoard[prettyIndex] = displayPiece(piece)
		}
	}

	for i := 0; i < (board.RowSize * board.ColumnSize); i++ {
		fmt.Print(prettyBoard[i])
		if ((i + 1) % board.RowSize) == 0 {
			fmt.Println("")
		}
	}
}

func displayPiece(piece Piece) string {
	switch piece.Shape {
	case "l":
		return "L"
	case "dot":
		return "O"
	case "square":
		return "X"
	case "line":
		return "|"
	}
	return ""
}
