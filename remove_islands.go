package main

import (
	"fmt"
)

func removeIslands(){

	// Given a matrix with 1's and 0's remove any island that's not
	// connected to the border
	land := [][]int{
		[]int{1, 0, 0, 0, 0, 0},
		[]int{0, 1, 0, 1, 1, 1},
		[]int{0, 0, 1, 0, 1, 0},
		[]int{1, 1, 0, 0, 1, 0},
		[]int{1, 0, 1, 1, 0, 0},
		[]int{1, 0, 0, 0, 0, 1},
	}

	adjList := [][2]int{}

	// looop over the graph to get the adjacency list
	for i := 0; i < len(land); i++ {
		if land[i][0] != 0 {
			adjList = append(adjList, [2]int{i, 0})
		}

		if land[i][len(land[i])-1] != 0 {
			adjList = append(adjList, [2]int{i, len(land[i])-1})
		}

		if i == 0 || i == len(land){
			for j := 0; j < len(land[i]); j++ {
				if land[i][j] != 0 {
					adjList = append(adjList, [2]int{i,j})
				}
			}
		}
	}

	for _, adjacency := range adjList {
		landDFS(land, adjacency[0], adjacency[1], 1)
	}

	for i := 0; i < len(land); i++ {
		for j := 0; j < len(land[i]); j++ {
			if land[i][j] > 1 {
				land[i][j] = 1
			} else {
				land[i][j] = 0
			}
			fmt.Printf("%d ", land[i][j])
		}
		fmt.Printf("\n")
	}
}

func landDFS(land [][]int, i, j, level int) {
	if land[i][j] == 0  || land[i][j] > 1{
		return
	}

	land[i][j] += level

	// Check up
	if i > 0 && land[i-1][j] != 0 {
		landDFS(land, i - 1, j, level + 1)
	}

	// check down
	if i < len(land)-1 && land[i+1][j] != 0{
		landDFS(land, i + 1, j, level + 1)
	}

	// check left
	if j > 0 && land[i][j-1] != 0 {
		landDFS(land, i, j - 1, level + 1)
	}

	// check right
	if j < len(land[i])-1 && land[i][j+1] != 0{
		landDFS(land, i, j + 1, level + 1)
	}
}
