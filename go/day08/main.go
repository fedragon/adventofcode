package main

import (
	"bufio"
	"fmt"
	"os"
)

type Grid [][]int

func main() {
	f, err := os.Open("../data/day08")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	grid := buildGrid(bufio.NewScanner(f))

	fmt.Println("total visible trees", countVisible(grid))

	fmt.Println("max scenic score", findMaxScenicScore(grid))
}

func buildGrid(scanner *bufio.Scanner) Grid {
	var grid Grid

	for scanner.Scan() {
		line := scanner.Text()

		var row []int
		for _, r := range line {
			row = append(row, int(r-'0'))
		}

		grid = append(grid, row)
	}

	return grid
}

func countVisible(grid Grid) int {
	var count int
	for rowIndex, row := range grid {
		cols := len(row)

		if rowIndex == 0 || rowIndex == len(grid)-1 {
			count += cols
			continue
		}

		for colIndex := range row {
			if colIndex == 0 || colIndex == cols-1 {
				count++
				continue
			}

			if isVisible(grid, rowIndex, colIndex) {
				count++
			}
		}
	}

	return count
}

func isVisible(grid Grid, row, col int) bool {
	tree := grid[row][col]

	top := true
	for r := 0; r < row; r++ {
		if grid[r][col] >= tree {
			top = false
		}
	}

	bottom := true
	for r := row + 1; r < len(grid); r++ {
		if grid[r][col] >= tree {
			bottom = false
		}
	}

	left := true
	for c := 0; c < col; c++ {
		if grid[row][c] >= tree {
			left = false
		}
	}

	right := true
	for c := col + 1; c < len(grid[row]); c++ {
		if grid[row][c] >= tree {
			right = false
		}
	}

	return top || bottom || left || right
}

func findMaxScenicScore(grid Grid) int {
	var max int
	for rowIndex, row := range grid {
		for colIndex := range row {
			if score := scenicScoreOf(grid, rowIndex, colIndex); score > max {
				max = score
			}
		}
	}

	return max
}

func scenicScoreOf(grid Grid, row, col int) int {
	tree := grid[row][col]

	if row == 0 || col == 0 || row == len(grid)-1 || col == len(grid[row])-1 {
		return 0
	}

	top := 0
	for r := row - 1; r >= 0; r-- {
		top++

		if grid[r][col] >= tree {
			break
		}
	}

	bottom := 0
	for r := row + 1; r < len(grid); r++ {
		bottom++

		if grid[r][col] >= tree {
			break
		}
	}

	left := 0
	for c := col - 1; c >= 0; c-- {
		left++

		if grid[row][c] >= tree {
			break
		}
	}

	right := 0
	for c := col + 1; c < len(grid[row]); c++ {
		right++

		if grid[row][c] >= tree {
			break
		}
	}

	return top * bottom * left * right
}
