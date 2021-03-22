package matrix

// If needed, you may define helper functions here.

// AreAdjacent returns true iff a and b are adjacent in lst.
func AreAdjacent(lst []int, a, b int) bool {
	indexA := 0
	for i, n := range lst {
		if n == a {
			indexA = i
		}
	}

	indexB := 0
	for j, m:= range lst {
		if m == b {
			indexB = j
		}
	}

	if indexA - 1 == indexB {
		return true
	} else if indexA + 1 == indexB {
		return true
	} else {
		return false
	}
}

// Transpose returns the transpose of the 2D matrix mat.
func Transpose(mat [][]int) [][]int {
	if len(mat) == 0 {
		return mat
	}

	arr := make([][]int, len(mat[0]))
	for i := range arr {
		arr[i] = make([]int, len(mat))
	}
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			arr[j][i] = mat[i][j]
		}
	}
	return arr
}

// AreNeighbors returns true iff a and b are Manhattan neighbors in the 2D
// matrix mat.
func AreNeighbors(mat [][]int, a, b int) bool {
	if len(mat) == 0 {
		return false
	}

	indexA := 0
	indexB := 0
	for i, n := range mat {
		if n[i] == a {
			indexA = i
		}
		for j, m := range n {
			if m == b {
				indexB = j
			}
		}
	}

	if indexA + 1 == indexB {
		return true
	} else if indexA - 1 == indexB {
		return true
	} else if indexA == indexB {
		return true
	} else {
		return false
	}
}
