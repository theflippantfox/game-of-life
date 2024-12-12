package main

import "fmt"

type BB uint64

func getBit(grid BB, square int) BB {
	return grid & (BB(1) << BB(square))
}

func setBit(grid *BB, square int) {
	mask := BB(1) << square
	*grid |= mask
}

func clearBit(grid *BB, square int) {
	mask := BB(1) << square
	*grid &^= mask
}

func toggleBit(grid *BB, square int) {
	mask := BB(1) << square
	*grid ^= mask
}

func printGrid(grid BB) {
	fmt.Println("Bitboard Grid:")
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			square := row*8 + col
			if getBit(grid, square) != 0 {
				fmt.Print("1 ")
			} else {
				fmt.Print("0 ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func initializeGrid(pattern []int) BB {
	grid := BB(0)
	for _, square := range pattern {
		setBit(&grid, square)
	}
	return grid
}

func startGame(grid *BB, generations int) {
	for gen := 0; gen < generations; gen++ {
		nextGrid := BB(0) // New grid to store the next generation

		for square := 0; square < 64; square++ {
			// Count neighbors for the current cell
			neighbors := countNeighbors(*grid, square)

			// Apply the rules of Game of Life
			if getBit(*grid, square) != 0 { // Cell is alive
				if neighbors == 2 || neighbors == 3 {
					setBit(&nextGrid, square) // Survives
				}
			} else { // Cell is dead
				if neighbors == 3 {
					setBit(&nextGrid, square) // Becomes alive
				}
			}
		}

		// Update the grid for the next generation
		*grid = nextGrid
		printGrid(*grid) // Optional: Print grid after each generation
	}
}

func countNeighbors(grid BB, square int) int {
	row, col := square/8, square%8 // Convert square index to row and col
	neighbors := 0

	// Iterate through the 8 possible neighbors
	for dr := -1; dr <= 1; dr++ {
		for dc := -1; dc <= 1; dc++ {
			if dr == 0 && dc == 0 {
				continue // Skip the current cell itself
			}
			r, c := row+dr, col+dc
			if r >= 0 && r < 8 && c >= 0 && c < 8 { // Ensure within bounds
				neighborSquare := r*8 + c
				if getBit(grid, neighborSquare) != 0 {
					neighbors++
				}
			}
		}
	}
	return neighbors
}

func main() {
	grid := initializeGrid([]int{2, 3, 18, 26, 22, 32, 33})
	printGrid(grid)

	startGame(&grid, 100)
}
