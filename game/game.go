package game

import (
	"fmt"

	"github.com/mitchellh/hashstructure"
)

func Solve(board Board) {
	boards := make([]Board, 1)
	allBoardHashes := make([]uint64, 1)
	boards[0] = board
	firstHash, _ := hashstructure.Hash(board, nil)
	allBoardHashes[0] = firstHash
	solving := true
	var solution Board

	for solving {
		nextBoards := make([]Board, 0)

		// pipe := emitBoards(boards)
		// results := processBoards(pipe, allBoardHashes)

		for i := 0; i < len(boards); i++ {
			var newBoards []Board
			newBoards, allBoardHashes = makeMoves(boards[i], allBoardHashes)
			nextBoards = append(nextBoards, newBoards...)

			// newBoards := <-results
			// nextBoards = append(nextBoards, newBoards...)
			// for board := range newBoards {
			// 	hash, _ := hashstructure.Hash(board, nil)
			// 	allBoardHashes = append(allBoardHashes, hash)
			// }
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
			fmt.Printf("Layer %v has %v board(s)\n", boards[0].Layer, len(boards))

			// if nextBoards[0].Layer <= 5 {
			// 	for _, board := range boards {
			// 		PrintBoard(board)
			// 		fmt.Println("")
			// 	}
			// }
		} else {
			fmt.Println("No more boards :(")
			return
		}
	}

	fmt.Println("Solved it!")
	PrintBoard(solution)
}

// func emitBoards(boards []Board) <-chan Board {
// 	out := make(chan Board)
// 	go func() {
// 		for _, board := range boards {
// 			out <- board
// 		}
// 		close(out)
// 	}()
// 	return out
// }

// func processBoards(in <-chan Board, allBoardHashes []uint64) <-chan []Board {
// 	out := make(chan []Board)
// 	go func() {
// 		for board := range in {
// 			out <- makeMoves(board, allBoardHashes)
// 		}
// 		close(out)
// 	}()
// 	return out
// }
