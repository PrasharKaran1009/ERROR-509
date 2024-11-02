package main

import (
	"fmt"
	"math"
)

func solution() int {
	var n int
	fmt.Scan(&n)
	
	// Initialize a 2D slice 'num' with all elements set to 0
	num := make([][]int, n)
	for i := 0; i < n; i++ {
		num[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&num[i][j])
		}
	}

	sum := 0 // read at last of code
	for i := 0; i < 2*n-1; i++ {
		mini := 0
		j, k := 0, 0

		// Calculate starting indices for diagonals
		if i < n {
			k = (n - 1) - i
		} else {
			j = (i - n) + 1
		}

		// Traverse along the diagonal
		for j < n && k < n {
			mini = int(math.Min(float64(mini), float64(num[j][k])))
			j++
			k++
		}

		sum += int(math.Abs(float64(mini))) // add absolute value of minimum element to sum
	}

	return sum
}

func main() {
	var t int
	fmt.Scan(&t)
	for t > 0 {
		t--
		fmt.Println(solution())
	}
}

/*
   04
   03 14
   02 13 24
   01 12 23 34
   00 11 22 33 44
   10 21 32 43
   20 31 42
   30 41
   40

Take minimum element along each diagonal, add its absolute value, and return the sum.
*/
