package game

import (
	"fmt"
	"time"
)

func Solve(board Board) {
	boards := make([]Board, 1)
	allBoardHashes := make(map[string]int, 1)
	boards[0] = board
	firstHash := hashBoard(board)
	allBoardHashes[firstHash] = 1
	solving := true
	var solution Board
	start := time.Now()
	fmt.Println("Starting board:")
	PrintBoard(board)

	for solving {
		nextBoards := make([]Board, 0)

		for i := 0; i < len(boards); i++ {
			var newBoards []Board
			newBoards, allBoardHashes = makeMoves(boards[i], allBoardHashes)
			nextBoards = append(nextBoards, newBoards...)
		}

		for _, newBoard := range nextBoards {
			if isSolution(newBoard) {
				fmt.Println("Found solution!!!!")
				solving = false
				solution = newBoard
			}
		}

		boards = nextBoards
		if len(boards) > 0 {
			t := time.Now()
			fmt.Printf("Layer %v has %v board(s) - %v\n", boards[0].Layer, len(boards), t.Sub(start))
		} else {
			fmt.Println("No more boards :(")
			return
		}
	}

	fmt.Println("Solved it!")
	PrintBoard(solution)
}
