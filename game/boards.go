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
	Id          int
}

type Board struct {
	Pieces     []Piece
	RowSize    int
	ColumnSize int
	Layer      int `hash:"ignore"`
}

var TenBoard = Board{
	Layer:      0,
	RowSize:    4,
	ColumnSize: 6,
	Pieces: []Piece{
		Piece{
			Id:    1,
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
			Id:    2,
			Shape: "dot",
			Coordinates: []Coordinate{
				Coordinate{
					X: 2,
					Y: 2,
				},
			},
		},
		Piece{
			Id:    3,
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
		Piece{
			Id:    2,
			Shape: "dot",
			Coordinates: []Coordinate{
				Coordinate{
					X: 0,
					Y: 3,
				},
			},
		},
		Piece{
			Id:    5,
			Shape: "square",
			Coordinates: []Coordinate{
				Coordinate{
					X: 1,
					Y: 3,
				},
				Coordinate{
					X: 2,
					Y: 3,
				},
				Coordinate{
					X: 1,
					Y: 4,
				},
				Coordinate{
					X: 2,
					Y: 4,
				},
			},
		},
		Piece{
			Id:    2,
			Shape: "dot",
			Coordinates: []Coordinate{
				Coordinate{
					X: 3,
					Y: 3,
				},
			},
		},
		Piece{
			Id:    7,
			Shape: "line",
			Coordinates: []Coordinate{
				Coordinate{
					X: 0,
					Y: 4,
				},
				Coordinate{
					X: 0,
					Y: 5,
				},
			},
		},
		Piece{
			Id:    2,
			Shape: "dot",
			Coordinates: []Coordinate{
				Coordinate{
					X: 1,
					Y: 5,
				},
			},
		},
		Piece{
			Id:    9,
			Shape: "l",
			Coordinates: []Coordinate{
				Coordinate{
					X: 3,
					Y: 4,
				},
				Coordinate{
					X: 3,
					Y: 5,
				},
				Coordinate{
					X: 2,
					Y: 5,
				},
			},
		},
	},
}

func cloneBoard(board Board) Board {
	newBoard := Board{}
	newBoard.RowSize = board.RowSize
	newBoard.ColumnSize = board.ColumnSize
	newBoard.Layer = board.Layer
	newBoard.Pieces = make([]Piece, 0)

	for i := 0; i < len(board.Pieces); i++ {
		piece := board.Pieces[i]
		newPiece := Piece{}

		newPiece.Shape = piece.Shape
		newPiece.Id = piece.Id
		newPiece.Coordinates = make([]Coordinate, 0)

		for i2 := 0; i2 < len(piece.Coordinates); i2++ {
			coordinate := piece.Coordinates[i2]
			newCoordinate := Coordinate{}

			newCoordinate.X = coordinate.X
			newCoordinate.Y = coordinate.Y

			newPiece.Coordinates = append(newPiece.Coordinates, newCoordinate)
		}

		newBoard.Pieces = append(newBoard.Pieces, newPiece)
	}

	return newBoard
}

func PrintBoard(board Board) {
	prettyBoard := make([]string, board.RowSize*board.ColumnSize)

	for i := 0; i < (board.RowSize * board.ColumnSize); i++ {
		prettyBoard[i] = " "
	}
	prettyBoard[0] = "@"
	prettyBoard[3] = "@"

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
