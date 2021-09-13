package main

// "Given a matrix of N rows and M columns. From m[i][j],
// we can move to m[i+1][j], if m[i+1][j] > m[i][j], or can
// move to m[i][j+1] if m[i][j+1] > m[i][j]. The task is print
// longest path length if we start from (0, 0)."

func longestMatrixPath(matrix [][]int) int {

	return longestMatrixPathDFS(matrix,0,0,1)
}

func longestMatrixPathDFS(matrix [][]int, x, y, currentLongestPath int) int {
	currentSum := matrix[x][y]

	// check up
	if y != 0 && matrix[x][y-1] > currentSum {
		return longestMatrixPathDFS(matrix, x, y-1, currentLongestPath+1)
	}

	// check down
	if y != len(matrix[x])-1 && matrix[x][y+1] > currentSum {
		return longestMatrixPathDFS(matrix, x, y + 1, currentLongestPath+1)
	}

	// check left
	if x != 0 && matrix[x-1][y] > currentSum {
		return longestMatrixPathDFS(matrix, x - 1, y, currentLongestPath+1)
	}

	// check right
	if x != len(matrix)-1 && matrix[x+1][y] > currentSum {
		return longestMatrixPathDFS(matrix, x + 1, y, currentLongestPath+1)
	}

	return currentLongestPath
}