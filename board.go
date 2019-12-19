package main

import (
	"fmt"
	"math/rand"
)

type Board struct {
	board           [8][8]int
	heuristicsBoard [8][8]int
	stateChanges    int
	restarts        int
	lowerNeighbors  int
}

func NewBoard() *Board {
	var board Board
	for i := 0; i < 8;i++ {
		board.board[i][rand.Intn(7)] = 1
	}
	board.GetNeighborStates()
	return &board
}

func (b *Board) PrintBoard() {
	for r := 0; r < 8; r++ {
		for c := 0; c < 8; c++ {
			fmt.Printf("%d", b.board[c][r])
			if c < 7 {
				fmt.Print(",")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func (b *Board) ResetQueens() {
	for i := 0; i < 8;i++ {
		for j := 0; j < 8; j++ {
			b.board[i][j] = 0
		}
	}
	for i := 0; i < 8;i++ {
		b.board[i][rand.Intn(7)] = 1
	}
	b.GetNeighborStates()
}

func (b *Board) GetRestarts() int {
	return b.restarts
}

func (b *Board) GetStateChanges() int {
	return b.stateChanges
}

func (b *Board) GetCurrentHeuristic() int {
	return b.checkDiagonals() + b.checkColumns() + b.checkRows()
}

func (b *Board) GetNeighborStates() {
	for i := 0; i < 8;i++ {
		for j := 0; j < 8; j++ {
			b.heuristicsBoard[i][j] = 0
		}
	}
	copyBoard := copy8DimArray(b.board)
	var boardHolder [8][8]int

	for c := 0; c < 8; c++ {
		for r := 0; r < 8; r++ {
			copyBoard[c][r] = 0
		}
		for r := 0; r < 8; r++ {
			copyBoard[c][r] = 1
			boardHolder = copy8DimArray(b.board)
			b.board = copy8DimArray(copyBoard)
			b.heuristicsBoard[c][r] = b.GetCurrentHeuristic()
			b.board = copy8DimArray(boardHolder)
			copyBoard[c][r] = 0
		}
		copyBoard = copy8DimArray(b.board)
	}
}

func (b *Board) CountLowerNeighbors() int {
	b.lowerNeighbors = 0
	currentH := b.GetCurrentHeuristic()
	for r := 0; r < 8; r++ {
		for c := 0; c < 8; c++ {
			if b.heuristicsBoard[c][r] < currentH {
				b.lowerNeighbors++
			}
		}
	}
	return b.lowerNeighbors
}

func (b *Board) ChangeState() {
	newC := -1
	newR := -1
	currentH := b.GetCurrentHeuristic()
	minH := currentH
	for r := 0; r < 8; r++ {
		for c := 0; c < 8; c++ {
			if b.heuristicsBoard[c][r] < minH {
				newC = c
				newR = r
				minH = b.heuristicsBoard[c][r]
			}
		}
	}
	for r := 0; r < 8; r++ {
		b.board[newC][r] = 0
	}
	b.board[newC][newR] = 1
	b.stateChanges++
	b.GetNeighborStates()
}

func (b *Board) checkDiagonals() int {
	h := 0

	for k := 0; k<8*2-1; k++ {
		count := 0
		for c := 0; c <= k; c++ {
			r := k-c
			if r<8 && c<8 {
				if b.board[c][r] == 1 {
					count++
				}
			}
		}
		if count > 1 {
			problems := (count*(count-1))/2
			h = h + problems
		}
	}

	for k := -7; k<8; k++ {
		count := 0
		for c := 0; c<=k+7; c++ {
			r := c-k
			if r<8 && c<8 && r>-1 && c >-1 {
				if b.board[c][r] == 1 {
					count++
				}
			}
		}
		if count > 1 {
			problems := (count*(count-1))/2
			h = h + problems
		}
	}

	return h
}

func (b *Board) checkRows() int {
	h := 0
	for r := 0; r < 8; r++ {
		count := 0
		for c := 0; c < 8; c++ {
			if b.board[c][r] == 1 {
				count++
			}
		}
		if count > 1 {
			problems := (count*(count-1))/2
			h = h + problems
		}
	}
	return h
}

func (b *Board) checkColumns() int {
	h := 0
	for c := 0; c < 8; c++ {
		count := 0
		for r := 0; r < 8; r++ {
			if b.board[c][r] == 1 {
				count++
			}
		}
		if count > 1 {
			problems := (count*(count-1))/2
			h= h + problems
		}
	}
	return h
}

func copy8DimArray(arr [8][8]int) [8][8]int {
	var copy [8][8]int
	for i := 0; i < 8;i++ {
		for j := 0; j < 8; j++ {
			copy[i][j] = arr[i][j]
		}
	}
	return copy
}

func (b *Board) IsSolved() bool {
	return b.GetCurrentHeuristic() == 0
}

func (b *Board) NeedsRestart() bool {
	return b.GetCurrentHeuristic() > 0 && b.CountLowerNeighbors() == 0
}

func RunLab() {
	board := NewBoard()
	for {
		fmt.Printf("Current h: %d\n", board.GetCurrentHeuristic())
		fmt.Println("Current state")
		board.PrintBoard()
		fmt.Printf("Neighbors found with lower h: %d\n", board.CountLowerNeighbors())
		if (board.NeedsRestart()) {
			fmt.Println("Restart")
			board.ResetQueens()
		} else {
			fmt.Println("Setting new current state")
			board.ChangeState()
		}

		fmt.Print("\n")
		if board.IsSolved() {
			break
		}
	}

	fmt.Println("Current state")
	board.PrintBoard()
	fmt.Println("Solution found")
	fmt.Printf("State changes: %d \n", board.GetStateChanges())
	fmt.Printf("Restarts: %d \n", board.GetRestarts())
}