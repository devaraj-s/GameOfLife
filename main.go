package main

import (
	"fmt"
)

type GameOfLife struct {
	Matrix [][]int
}

type CellDetails struct {
	CurrentGenerationCellIndex [2]int
	CurrentGeneration          int
	NewGeneration              int
	Neighbours                 [3][3]int
	TotalNoOfLiveNeighbour     int
	TotalNoOfDeadNeighbour     int
}

func main() {
	fmt.Println("Game of life")
	data := [][]int{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}

	newGeneration := [][]int{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
	gameOfLife := &GameOfLife{data}

	for i, row := range data {
		for j, _ := range row {
			cellIndex := [2]int{i, j}
			cell := gameOfLife.getNeighbours(cellIndex)
			cell.Transition()

			//fmt.Printf("Cell details %v -> %+v\n", data[i][j], cell)
			newGeneration[i][j] = cell.NewGeneration

		}
		fmt.Println()
	}
	//fmt.Println(data)
	fmt.Println(newGeneration)
}

func (g *GameOfLife) getNeighbours(cell [2]int) *CellDetails {

	rowIndex := cell[0]
	columnIndex := cell[1]
	var neighbours [3][3]int
	countOfLiveCells := 0
	countOfDeadCells := 0
	neighboursRowIndex := 0
	for i := (rowIndex - 1); i <= rowIndex+1; i++ {
		neighboursColIndex := 0
		for j := (columnIndex - 1); j <= columnIndex+1; j++ {
			if i == -1 || i == len(g.Matrix) {
				neighbours[neighboursRowIndex] = [3]int{-1, -1, -1}
				continue
			}

			if j >= len(g.Matrix) || j == -1 {
				neighbours[neighboursRowIndex][neighboursColIndex] = -1
			} else {
				cell := g.Matrix[i][j]
				neighbours[neighboursRowIndex][neighboursColIndex] = cell

				//ignoring the current cell
				if i == rowIndex && j == columnIndex {
					continue
				}
				if cell == 0 {
					countOfDeadCells++
				} else {
					countOfLiveCells++
				}
			}
			neighboursColIndex++
		}
		neighboursRowIndex++
	}

	neighbourDetails := &CellDetails{Neighbours: neighbours,
		TotalNoOfDeadNeighbour: countOfDeadCells,
		TotalNoOfLiveNeighbour: countOfLiveCells,
		CurrentGeneration:      g.Matrix[cell[0]][cell[1]], CurrentGenerationCellIndex: cell}
	return neighbourDetails
}

//Any live cell with fewer than two live neighbours dies, as if by underpopulation.
func (cell *CellDetails) Transition() {

	//if current generation is live
	if cell.CurrentGeneration == 1 {

		//Any live cell with fewer than two live neighbours dies, as if by underpopulation.
		//Any live cell with two or three live neighbours lives on to the next generation.
		//Any live cell with more than three live neighbours dies, as if by overpopulation.
		if cell.TotalNoOfLiveNeighbour < 2 {
			cell.NewGeneration = 0
		} else if cell.TotalNoOfLiveNeighbour >= 2 && cell.TotalNoOfLiveNeighbour <= 3 {
			cell.NewGeneration = 1
		} else if cell.TotalNoOfLiveNeighbour > 3 {
			cell.NewGeneration = 0
		}

	} else {

		//Any dead cell with three live neighbours becomes a live cell, as if by reproduction.
		if cell.TotalNoOfLiveNeighbour == 3 {
			cell.NewGeneration = 1
		}
	}
}
